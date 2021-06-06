package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat04/deadlines/proto"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := proto.NewHelloClient(conn)
	dealine := time.Now().Add(time.Duration(2 * time.Second))
	ctx, cancel := context.WithDeadline(context.Background(), dealine)
	defer func() {
		log.Print("context cancel ")
		cancel()
	}()

	//cancel 错误
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	log.Printf("run cancel test ...")
	//	cancel()
	//}()

	res, err := client.Echo(ctx, &proto.String{"abcdef"})
	if err != nil {
		log.Printf("echo error: %v ", status.Code(err))
	} else {
		log.Print("res value: ", res.Str)
	}

}
