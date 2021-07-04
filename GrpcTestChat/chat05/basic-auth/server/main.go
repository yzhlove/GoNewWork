package main

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat05/basic-auth/echo"
	"log"
	"net"
	"strings"
)

var (
	port            = ":50051"
	errMissMetadata = status.Errorf(codes.InvalidArgument, "miss metadata")
	errInvalidToken = status.Errorf(codes.Unauthenticated, "invalid credentials")
	crtFile         = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/GrpcTestChat/chat05/secure-channel/certs/server.crt"
	keyFile         = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/GrpcTestChat/chat05/secure-channel/certs/server.key"
)

func main() {

	cert, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.UnaryInterceptor(checkIntercept),
	}
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

func checkIntercept(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, h grpc.UnaryHandler) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissMetadata
	}

	if !check(md["authorization"]) {
		return nil, errInvalidToken
	}

	return h(ctx, req)
}

func check(auths []string) bool {
	fmt.Println("auths => ", auths)
	if len(auths) > 0 {
		token := strings.TrimPrefix(auths[0], "Basic ")
		dec, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("decoder ===> ", string(dec))
		return token == base64.StdEncoding.EncodeToString([]byte("root:root"))
	}
	return false
}
