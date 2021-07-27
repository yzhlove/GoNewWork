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
	crtFile  = "/Users/yostar/workSpace/GoNewWork/GrpcTestChat/chat05/jwtauth/certs/server.crt"
)

var secret = []byte("*#06#*")

func main() {

	creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
	if err != nil {
		log.Fatal(err)
	}
	auth := basicAuth{username: "root", password: "root"}
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(auth),
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cc := echo.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := cc.Echo(ctx, &echo.String{Value: "what are you doing?"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("resp value: %s \n", resp.Value)
}

type basic struct {
	UserName string
	Password string
	*jwt.StandardClaims
}

func obtained(username, password string) *basic {
	return &basic{
		UserName:       username,
		Password:       password,
		StandardClaims: &jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour).Unix()},
	}
}

type basicAuth struct {
	username, password string
}

func (b basicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, obtained(b.username, b.password))
	//secret type must []byte
	enc, err := token.SignedString(secret)
	if err != nil {
		log.Fatal("GetRequestMetadata Error:" + err.Error())
	}
	return map[string]string{"x-token": enc}, nil
}

func (b basicAuth) RequireTransportSecurity() bool {
	return true
}
