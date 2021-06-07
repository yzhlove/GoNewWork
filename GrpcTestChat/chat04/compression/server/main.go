package main

import (
	"context"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip" // 服务端注册压缩器
	"grpc-test-chat/chat04/compression/proto"
	"log"
	"net"
	"strings"
)

func main() {

	if l, err := net.Listen("tcp", ":50051"); err != nil {
		log.Fatal(err)
	} else {
		s := grpc.NewServer()
		proto.RegisterHelloServer(s, &server{})
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}

}

type server struct{}

func (s *server) Echo(ctx context.Context, str *proto.String) (*proto.String, error) {
	return &proto.String{Str: strings.ToUpper(str.Str)}, nil
}
