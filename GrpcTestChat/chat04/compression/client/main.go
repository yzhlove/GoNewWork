package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
	"grpc-test-chat/chat04/compression/proto"
	"log"
)

func main() {
	testUseCompress1()
	testUseCompress2()
}

//客户端注册压缩器的方式1
func testUseCompress1() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloClient(conn)

	res, err := client.Echo(context.Background(), &proto.String{Str: "love"}, grpc.UseCompressor(gzip.Name))
	if err != nil {
		log.Fatal(err)
	}

	log.Print("echo :", res.Str)
}

//客户端注册压缩器的方式2
func testUseCompress2() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloClient(conn)

	res, err := client.Echo(context.Background(), &proto.String{Str: "love"}, grpc.UseCompressor(gzip.Name))
	if err != nil {
		log.Fatal(err)
	}

	log.Print("echo :", res.Str)

}
