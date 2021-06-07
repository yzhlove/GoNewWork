package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test-chat/chat04/mulitplex/proto"
	"log"
	"net"
	"strings"
)

func main() {

	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	proto.RegisterAServer(s, &A{})
	proto.RegisterBServer(s, &B{})
	//proto.RegisterAServer(s, &C{})
	//proto.RegisterBServer(s, &C{})

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

type A struct{}

func (a *A) Echo(ctx context.Context, str *proto.String) (*proto.String, error) {
	return &proto.String{Str: "A" + strings.ToUpper(str.Str)}, nil
}

type B struct{}

func (b *B) Echo(ctx context.Context, str *proto.String) (*proto.String, error) {
	return &proto.String{Str: "B" + strings.ToLower(str.Str)}, nil
}

//type C struct{}
//
//func (c *C) Echo(ctx context.Context, str *proto.String) (*proto.String, error) {
//	return &proto.String{Str: "C" + str.Str}, nil
//}
