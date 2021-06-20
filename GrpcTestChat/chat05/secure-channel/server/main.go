package main

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-test-chat/chat05/secure-channel/echo"
	"log"
	"net"
	"strings"
)

var (
	port    = ":50051"
	crtFile = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/GrpcTestChat/chat05/secure-channel/certs/server.crt"
	keyFile = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/GrpcTestChat/chat05/secure-channel/certs/server.key"
)

func main() {

	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{grpc.Creds(credentials.NewServerTLSFromCert(&cert))}
	s := grpc.NewServer(opts...)
	echo.RegisterHelloServer(s, &server{})

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

type server struct{}

func (s *server) Echo(ctx context.Context, str *echo.String) (*echo.String, error) {
	return &echo.String{Value: strings.ToUpper(str.Value)}, nil
}
