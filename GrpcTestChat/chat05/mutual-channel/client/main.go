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
	"time"
)

var (
	address  = "localhost:50051"
	hostname = "localhost"
	common   = "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/GrpcTestChat/chat05/mutual-channel/certs/"
	crtFile  = common + "client.crt"
	keyFile  = common + "client.key"
	caFile   = common + "ca.crt"
)

func main() {

	creds, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	cert := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}

	if ok := cert.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed append to ca certs")
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName:   hostname,
			Certificates: []tls.Certificate{creds},
			RootCAs:      cert,
		})),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cc := echo.NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := cc.Echo(ctx, &echo.String{Value: "abced"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("resp value -> ", resp.Value)

}
