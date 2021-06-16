package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat04/compression/proto"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cc := proto.NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := cc.Echo(ctx, &proto.String{Str: "anbced"})
	if err != nil {
		r := status.Code(err)
		log.Printf("error result ->  %v", r)
	} else {
		log.Println("result value -> ", resp.Str)
	}

}
