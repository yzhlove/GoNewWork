package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test-chat/chat04/compression/proto"
	"log"
	"net"
	"strings"
	"time"
)

func main() {

	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	proto.RegisterHelloServer(s, &server{})

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

type server struct{}

func (s *server) Echo(ctx context.Context, str *proto.String) (*proto.String, error) {
	time.Sleep(time.Second * 5)
	return &proto.String{Str: strings.ToUpper(str.Str)}, nil
}
