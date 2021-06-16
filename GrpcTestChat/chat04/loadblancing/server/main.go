package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	echo "grpc-test-chat/chat04/loadblancing/proto"
	"log"
	"net"
	"sync"
)

var address = []string{":50051", ":50052"}

func main() {

	var wg sync.WaitGroup
	for _, addr := range address {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			runEcho(addr)
		}(addr)
	}
	wg.Wait()
}

type server struct {
	addr string
}

func (s *server) Echo(ctx context.Context, str *echo.String) (*echo.String, error) {
	log.Println("ping address : ", s.addr)
	msg := &echo.String{Str: fmt.Sprintf("echo %s: addr:<%s>", str.Str, s.addr)}
	return msg, nil
}

func runEcho(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	echo.RegisterHelloServer(s, &server{addr: addr})
	log.Printf("serveing on addr:%s ", addr)
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
