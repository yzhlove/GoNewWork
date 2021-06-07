package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test-chat/chat04/mulitplex/proto"
	"log"
)

func main() {

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	a := proto.NewAClient(conn)
	b := proto.NewBClient(conn)

	res, err := a.Echo(context.Background(), &proto.String{Str: "abcde"})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("result -> ", res.Str)

	res, err = b.Echo(context.Background(), &proto.String{Str: "abcde"})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("result -> ", res.Str)

	c := proto.NewAClient(conn)
	res, err = c.Echo(context.Background(), &proto.String{Str: "abcde"})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("result -> ", res.Str)

	c = proto.NewBClient(conn)
	res, err = c.Echo(context.Background(), &proto.String{Str: "abcde"})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("result -> ", res.Str)

}
