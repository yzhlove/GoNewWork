package main

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-test-chat/chat05/jwtauth/echo"
	"log"
	"time"
)

const (
	address  = "localhost:50051"
	hostname = "localhost"
	certFile = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/GrpcTestChat/chat05/oauth/certs/server.crt"
	secret   = "*#06#*"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile(certFile, hostname)
	if err != nil {
		log.Fatal(err)
	}

	basic := Gener("yzh", "123456")
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(basic),
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cc := echo.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := cc.Echo(ctx, &echo.String{Str: "what are you doing?"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("resp value: %s \n", resp.Str)

}

type basicCheck struct {
	UserName string
	PassWord string
	*jwt.StandardClaims
}

func Gener(username, password string) basicCheck {
	return basicCheck{
		UserName: username,
		PassWord: password,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "love",
		},
	}
}

func GenerToken(basic basicCheck) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, basic)
	return token.SignedString(secret)
}

func (b basicCheck) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	token, err := GenerToken(b)
	if err != nil {
		return nil, err
	}

	return map[string]string{"x-token": token}, nil
}

func (b basicCheck) RequireTransportSecurity() bool {
	return true
}
