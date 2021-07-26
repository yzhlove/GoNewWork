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

const (
	port     = ":50051"
	certFile = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/GrpcTestChat/chat05/oauth/certs/server.crt"
	keyFile  = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/GrpcTestChat/chat05/oauth/certs/server.key"
	secret   = "*#06#*"
)

var (
	errMissMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken = status.Errorf(codes.Unauthenticated, "invalid token")
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

	server := grpc.NewServer(opts...)
	echo.RegisterEchoServiceServer(server, &ee{})
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listener on :%s\n", port)

	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

type ee struct{}

func (e *ee) Echo(ctx context.Context, str *echo.String) (*echo.String, error) {
	resp := &echo.String{Str: strings.ToUpper(str.Str)}
	return resp, nil
}

func check(strs []string) bool {
	if len(strs) > 0 {
		tokenstr := strings.Split(strs[0], " ")
		return tokenstr[len(tokenstr)-1] == secret
	}
	return false
}

func ensureToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handle grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissMetadata
	}
	log.Printf("data: %+v \n", md)
	if !check(md["authorization"]) {
		return nil, errInvalidToken
	}
	return handle(ctx, req)
}
