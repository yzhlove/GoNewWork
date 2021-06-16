package main

import (
	"context"
	"fmt"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat04/errorhandle/proto"
	"log"
	"net"
	"strings"
)

func main() {

	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	proto.RegisterHelloServer(s, &server{})

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}

type server struct{}

func (s *server) Echo(ctx context.Context, ss *proto.String) (*proto.String, error) {

	if ss.Str == "-1" {
		log.Print("error message.")

		errStatus := status.New(codes.InvalidArgument, "Invalid Arguments")

		ds, err := errStatus.WithDetails(&epb.BadRequest_FieldViolation{
			Field:       "ID",
			Description: fmt.Sprintf("message value out of range: %s", ss.Str),
		})

		if err != nil {
			return nil, errStatus.Err()
		}

		return nil, ds.Err()
	}

	return &proto.String{Str: strings.ToUpper(ss.Str)}, nil
}
