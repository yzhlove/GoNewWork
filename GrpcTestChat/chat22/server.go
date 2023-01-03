package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test-chat/chat22/echo"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", ":8887")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	echo.RegisterEchoServiceServer(server, &service{})
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}

type service struct{}

func (s *service) Echo(ctx context.Context, in *echo.Msg_Nil) (*echo.Msg_Nil, error) {
	return &echo.Msg_Nil{}, nil
}
