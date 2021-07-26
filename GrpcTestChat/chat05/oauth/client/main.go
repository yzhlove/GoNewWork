package main

import (
	"context"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"grpc-test-chat/chat05/oauth/echo"
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

	perRPC := oauth.NewOauthAccess(&oauth2.Token{AccessToken: secret})
	cerds, err := credentials.NewClientTLSFromFile(certFile, hostname)
	if err != nil {
		log.Fatal(err)
	}
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(perRPC),
		grpc.WithTransportCredentials(cerds),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cc := echo.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := cc.Echo(ctx, &echo.String{Str: "hello world"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("resp value: %s \n", resp.Str)

}
