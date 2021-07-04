package main

import (
	"context"
	"encoding/base64"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-test-chat/chat05/basic-auth/echo"
	"log"
	"time"
)

var (
	address = "localhost:50051"
	crtFile = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/GrpcTestChat/chat05/secure-channel/certs/server.crt"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile(crtFile, "localhost")
	if err != nil {
		log.Fatal(err)
	}

	auth := basicAuth{username: "root", password: "root"}
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(auth),
		grpc.WithTransportCredentials(creds),
	}

	cc, err := grpc.Dial(address, opts...)
	defer cc.Close()

	c := echo.NewHelloClient(cc)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := c.Echo(ctx, &echo.String{Value: "hello world"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("resp ", resp.Value)

}

type basicAuth struct {
	username, password string
}

func (b basicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	auth := b.username + ":" + b.password
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{"authorization": "Basic " + enc}, nil
}

func (b basicAuth) RequireTransportSecurity() bool {
	return true
}
