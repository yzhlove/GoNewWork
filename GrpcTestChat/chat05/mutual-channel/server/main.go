package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-test-chat/chat05/mutual-channel/echo"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

var (
	common  = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/GrpcTestChat/chat05/mutual-channel/certs/"
	port    = ":50051"
	crtFile = common + "server.crt"
	keyFile = common + "server.key"
	caFile  = common + "ca.crt"
)

func main() {

	//通过服务端证书和密钥直接创建x.509密钥对
	creds, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	//通过CA创建证书池
	cert := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}

	//将来自CA客户端的证书附加到证书池中
	if ok := cert.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append client certs")
	}

	opts := []grpc.ServerOption{
		//通过创建TLS凭证为所有传入的连接启用TLS
		grpc.Creds(credentials.NewTLS(&tls.Config{
			ClientAuth:   tls.RequireAndVerifyClientCert,
			Certificates: []tls.Certificate{creds},
			ClientCAs:    cert,
		})),
	}

	//通过传入TLS服务器凭证创建新的grpc服务器实例
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
