package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test-chat/chat09/echo"
	"log"
	"net"
)

type Req struct {
}

type Resp struct {
	resp interface{}
	err  error
}

type Manager struct {
	reqCh *Req
}

func (m *Manager) Run() error {

	return nil
}

type service struct {
}

func (s *service) Echo(ctx context.Context, req *echo.Req) (*echo.Resp, error) {

	return nil, nil
}

func (s *service) Intercept(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, invork grpc.UnaryHandler) (interface{}, error) {

	return nil, nil
}

const (
	_port = ":55566"
)

func main() {
	l, err := net.Listen("tcp", _port)
	if err != nil {
		log.Fatalf("net.listener error:%v", err)
	}

	service := &service{}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(service.Intercept))
	echo.RegisterEchoServiceServer(server, service)

	if err := server.Serve(l); err != nil {
		log.Fatalf("running grpc server error:%v", err)
	}

}
