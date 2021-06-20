package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-test-chat/chat05/secure-channel/echo"
	"log"
	"time"
)

var (
	address  = "localhost:50051"
	hostname = "localhost"
	crtFile  = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/GrpcTestChat/chat05/secure-channel/certs/server.crt"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
	if err != nil {
		log.Fatal("credentials:", err)
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	conn, err := grpc.Dial(address, opts...)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cc := echo.NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := cc.Echo(ctx, &echo.String{Value: "abcde"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("resp value => ", resp.Value)

}
