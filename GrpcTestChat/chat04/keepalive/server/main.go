package main

import (
	"context"
	"grpc-test-chat/chat04/compression/proto"
	"strings"
)

func main() {

}

type server struct{}

func (s *server) Echo(ctx context.Context, str *proto.String) (*proto.String, error) {
	return &proto.String{Str: strings.ToUpper(str.Str)}, nil
}
