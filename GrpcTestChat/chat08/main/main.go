package main

import (
	"context"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat08/echo"
	"log"
	"net"
	"time"
)

const (
	stuName   = "student"
	teachName = "teacher"
)

func main() {

	server()
	client()

}

type service struct{}

func (s *service) Echo(ctx context.Context, req *echo.Req) (*echo.Resp, error) {
	resp := &echo.Resp{Desc: "service:" + req.Type}
	switch req.Type {
	case stuName:
		student := &echo.Student{Name: "大宝", Subject: "数学", Address: "China Shanghai"}
		any, err := types.MarshalAny(student)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "student any error: %s", err)
		}
		resp.Character = any
	case teachName:
		teacher := &echo.Teacher{Name: "大壮", Project: "语文", Salary: 8000}
		any, err := types.MarshalAny(teacher)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "teacher any error: %s", err)
		}
		resp.Character = any
	default:
		return nil, status.Errorf(codes.Unknown, "not found type:%s", req.Type)
	}
	return resp, nil
}

func server() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	echo.RegisterHelloServiceServer(s, &service{})

	go func() {
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()
}

func client() {

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cc := echo.NewHelloServiceClient(conn)

	parse := func(resp *echo.Resp) {
		if resp == nil {
			log.Println("resp is nil")
			return
		}
		log.Printf("resp type:%s ", resp.Desc)

		if types.Is(resp.Character, &echo.Student{}) {
			student := &echo.Student{}
			if err := types.UnmarshalAny(resp.Character, student); err != nil {
				log.Fatal(err)
			}
			log.Println(" student -> ", student.Name, student.Subject, student.Address)
		}

		if types.Is(resp.Character, &echo.Teacher{}) {
			teacher := &echo.Teacher{}
			if err := types.UnmarshalAny(resp.Character, teacher); err != nil {
				log.Fatal(err)
			}
			log.Println(" teacher -> ", teacher.Name, teacher.Project, teacher.Salary)
		}

	}

	resp, err := cc.Echo(ctx, &echo.Req{Type: "abc"})
	if err != nil {
		log.Println(err)
	} else {
		parse(resp)
	}

	resp, err = cc.Echo(ctx, &echo.Req{Type: stuName})
	if err != nil {
		log.Println(err)
	} else {
		parse(resp)
	}

	resp, err = cc.Echo(ctx, &echo.Req{Type: teachName})
	if err != nil {
		log.Println(err)
	} else {
		parse(resp)
	}
}
