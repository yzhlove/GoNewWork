package main

import (
	"context"
	"crypto/tls"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat05/jwtauth/echo"
	"log"
	"net"
	"strings"
)

var (
	errMissData = status.Errorf(codes.InvalidArgument, "metadata data miss .")
	errToken    = status.Errorf(codes.Unauthenticated, "token is invalid")
)

const (
	certFile = "/Users/yostar/workSpace/GoNewWork/GrpcTestChat/chat05/jwtauth/certs/server.crt"
	keyFile  = "/Users/yostar/workSpace/GoNewWork/GrpcTestChat/chat05/jwtauth/certs/server.key"
)

var secret = []byte("*#06#*")

func main() {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.UnaryInterceptor(ensureToken),
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
	return &echo.String{Value: strings.ToUpper(str.Value)}, nil
}

func ensureToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handle grpc.UnaryHandler) (interface{}, error) {

	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissData
	}

	if !check(data["x-token"]) {
		return nil, errToken
	}

	return handle(ctx, req)
}

type agent struct {
	UserName string
	Password string
	*jwt.StandardClaims
}

func check(strs []string) bool {
	if len(strs) > 0 {
		log.Printf("token string slice: %v \n", strs)
		ant := &agent{}
		token, err := jwt.ParseWithClaims(strs[0], ant, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		log.Printf("agent: %+v \n", ant)

		if err != nil {
			log.Printf("token parse error: %v \n", err)
			return false
		}

		if err := token.Claims.Valid(); err != nil {
			log.Printf("token is invalid error: %v \n", err)
			return false
		}

		_, ok := token.Claims.(*agent)
		return ok && token.Valid
	}
	return false
}
