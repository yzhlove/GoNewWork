package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-test-chat/chat04/metadata/proto"
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

	data := metadata.Pairs("event", "client", "event", time.Now().Format(time.RFC3339))
	ctx := metadata.NewOutgoingContext(context.Background(), data)
	ctx = metadata.AppendToOutgoingContext(ctx, "k1", "v1", "k2", "v3", "k3", "v3")

	var header, trailer metadata.MD

	res, err := cc.Echo(ctx, &proto.Empty{}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("res.value => ", res.Str)

}

