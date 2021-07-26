package main

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat05/oauth/echo"
	"log"
	"net"
	"strings"
)

var (
	errMissData = status.Errorf(codes.InvalidArgument, "metadata data miss .")
	errToken    = status.Errorf(codes.Unauthenticated, "token is invalid")
)

const (
	certFile = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/GrpcTestChat/chat05/jwtauth/certs/server.crt"
	keyFile  = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/GrpcTestChat/chat05/jwtauth/certs/server.key"
)

func main() {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ensureToken),
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	service := grpc.NewServer(opts...)
	echo.RegisterEchoServiceServer(service, &server{})

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listener on the: localhost:50051")
	if err := service.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

type server struct{}

func (s *server) Echo(ctx context.Context, str *echo.String) (*echo.String, error) {
	return &echo.String{Str: strings.ToUpper(str.Str)}, nil
}

func ensureToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handle grpc.UnaryHandler) (interface{}, error) {

	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissData
	}

	log.Printf("x-token: %+v \n", data["x-token"])

	return handle(ctx, req)
}
