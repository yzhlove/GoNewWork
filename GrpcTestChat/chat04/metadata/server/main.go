package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-test-chat/chat04/metadata/proto"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func main() {

	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.MaxConcurrentStreams(1025))
	proto.RegisterHelloServer(s, &server{})

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}

type server struct{}

func (s *server) Echo(ctx context.Context, _ *proto.Empty) (*proto.String, error) {

	data, ok := metadata.FromIncomingContext(ctx)
	if ok {
		show(data)
	}

	var str string
	if r := data.Get("event"); len(r) > 0 {
		str = strings.Join(r, "-")
	} else {
		str = "noway"
	}

	header := metadata.New(map[string]string{"server": "echo", "stamp": "2021-06-18"})
	grpc.SendHeader(ctx, header)

	return &proto.String{Str: str}, nil
}

func (s *server) LoopEcho(st proto.Hello_LoopEchoServer) error {

	trailer := metadata.Pairs("time", time.Now().Format(time.RFC3339))
	st.SetTrailer(trailer)

	header := metadata.New(map[string]string{"local": "loop_echo", "ts": time.Now().Format(time.RFC3339)})
	st.SendHeader(header)

	for {
		if str, err := st.Recv(); err != nil {
			if err == io.EOF {
				log.Printf("loop echo break for loop.")
				break
			}
			log.Print("echo recv error:", err)
		} else {
			st.Send(&proto.String{Str: strings.ToUpper(str.Str)})
		}
	}

	return nil
}

func show(md metadata.MD) {
	for k, v := range md {
		for _, x := range v {
			log.Printf("k = %v v = %v", k, x)
		}
	}
}
