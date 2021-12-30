package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat09/echo"
	"log"
	"net"
)

func fetchMonitor(ctx context.Context) *Monitor {
	res := ctx.Value(monitorType{})
	if res != nil {
		if m, ok := res.(*Monitor); ok {
			return m
		}
	}
	return nil
}

type service struct{}

func (s *service) Echo(ctx context.Context, req *echo.Req) (*echo.Resp, error) {

	monitor := fetchMonitor(ctx)
	if monitor == nil {
		return nil, status.Error(codes.Aborted, "monitor type assert error")
	}
	monitor.Push(req.Msg)
	resp := &echo.Resp{Msg: fmt.Sprintf("%s:Ok", req.Msg), Code: 200}
	return resp, nil
}

func (s *service) Intercept(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, invork grpc.UnaryHandler) (interface{}, error) {

	monitor := New()
	defer monitor.Clear()

	ctx = context.WithValue(ctx, monitorType{}, monitor)
	resp, err := invork(ctx, req)
	if err != nil {
		return nil, err
	}

	monitor.Show()

	return resp, nil
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

	fmt.Println("listening on:", l.Addr())
	if err := server.Serve(l); err != nil {
		log.Fatalf("running grpc server error:%v", err)
	}

}
