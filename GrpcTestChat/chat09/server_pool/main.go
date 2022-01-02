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
	"strings"
)

type mtype struct{}

const _port = ":55566"

func main() {
	l, err := net.Listen("tcp", _port)
	if err != nil {
		log.Fatal(err)
	}

	service := &service{
		pool:    New(),
		manager: NewHandleManager(),
	}

	service.manager.run()

	server := grpc.NewServer(grpc.UnaryInterceptor(service.Intercept))
	echo.RegisterEchoServiceServer(server, service)

	log.Println("listener on ", l.Addr())
	if err := server.Serve(l); err != nil {
		log.Fatalf("running grpc server error:%v", err)
	}
}

type service struct {
	pool    *poolMonitor
	manager *handleManager
}

func fetch(ctx context.Context) *monitor {
	if monitor, ok := ctx.Value(mtype{}).(*monitor); ok {
		return monitor
	}
	return nil
}

func (s *service) Echo(ctx context.Context, req *echo.Req) (*echo.Resp, error) {

	fmt.Println("echo: handle request -> ", req)

	monitor := fetch(ctx)
	if monitor == nil {
		return nil, status.Error(codes.NotFound, "monitor is empty.")
	}

	monitor.Push(strings.Split(req.Msg, "-")...)
	resp := &echo.Resp{Msg: monitor.Merge(), Code: 200}
	return resp, nil
}

func (s *service) Intercept(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, invoke grpc.UnaryHandler) (interface{}, error) {

	monitor := s.pool.Get()
	defer s.pool.Put(monitor)
	log.Println("[server] ", info.FullMethod, " req:", req)
	ctx = context.WithValue(ctx, mtype{}, monitor)
	s.manager.DoReq(Req{method: info.FullMethod, ctx: ctx, request: req, handle: invoke})
	return s.manager.DoResp()
}
