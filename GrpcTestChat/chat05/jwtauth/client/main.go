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
	certFile = "/Users/yostar/workSpace/GoNewWork/GrpcTestChat/chat05/jwtauth/certs/server.crt"
	secret   = "*#06#*"
)

func main() {

	cerds, err := credentials.NewClientTLSFromFile(certFile, hostname)
	if err != nil {
		log.Fatal(err)
	}

	basic := obtained("yzh", "123456")
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(basic),
		grpc.WithTransportCredentials(cerds),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatal("dial error:" + err.Error())
	}
	defer conn.Close()

	cc := echo.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := cc.Echo(ctx, &echo.String{Value: "hello world"})
	if err != nil {
		log.Fatal("echo error:" + err.Error())
	}

	log.Printf("resp value: %s \n", resp.Value)

}

type basicAuth struct {
	UserName string
	PassWord string
	*jwt.StandardClaims
}

func obtained(username, password string) *basicAuth {
	return &basicAuth{
		UserName:       username,
		PassWord:       password,
		StandardClaims: &jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour).Unix()},
	}
}

func (b *basicAuth) RequireTransportSecurity() bool { return true }

func (b *basicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, b)
	tokenstr, err := token.SignedString(secret)
	log.Printf("token string: %s \n", tokenstr)
	return map[string]string{"x-token": tokenstr}, err
}
