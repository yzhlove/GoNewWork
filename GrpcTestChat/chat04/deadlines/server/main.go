package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test-chat/chat04/deadlines/proto"
	"log"
	"net"
)

func main() {

	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	proto.RegisterHelloServer(s, &server{})

	if err := s.Serve(l); err != nil {
		log.Fatalln(err)
	}
}

type server struct{}

//func (s *server) Echo(ctx context.Context, str *proto.String) (*proto.String, error) {
//
//	sleep := 5
//	time.Sleep(time.Duration(sleep) * time.Second)
//
//	if ctx.Err() == context.DeadlineExceeded {
//		log.Printf("RPC has reached deadline exceeded state: %s ", ctx.Err())
//		return nil, ctx.Err()
//	}
//
//	if ctx.Err() == context.Canceled {
//		log.Printf("Rpc has reached cancel state: %s ", ctx.Err())
//	}
//
//	log.Println("Echo value-> ", str.Str)
//	return &proto.String{Str: strings.ToUpper(str.Str)}, nil
//}

func (s *server) Echo(ctx context.Context, str *proto.String) (*proto.String, error) {

	select {
	case <-ctx.Done():
		log.Printf("context done -> %v ", ctx.Err())
	}
	//proto.string -> return nil is panic
	return &proto.String{}, nil
}
