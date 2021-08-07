package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat06/contions/echo"
	"log"
	"net"
	"time"
)

func main() {

	go server()
	peer()
	time.Sleep(time.Second)
}

type service struct{}

func (s *service) Echo(ctx context.Context, req *echo.EchoReq) (*echo.EchoResp, error) {
	resp := &echo.EchoResp{Desc: "Author:" + req.Name}
	switch req.Name {
	case "student":
		stu := &echo.Student{Name: "金大壮", Class: "数学"}
		if any, err := ptypes.MarshalAny(stu); err != nil {
			return nil, status.Errorf(codes.Unknown, "student marshal error:%s", err)
		} else {
			resp.Character = any
		}
	case "teacher":
		teacher := &echo.Teacher{Name: "李二狗子", Project: "语文", Salary: "8000"}
		if any, err := ptypes.MarshalAny(teacher); err != nil {
			return nil, status.Errorf(codes.Unknown, "teacher marsh error:%s", err)
		} else {
			resp.Character = any
		}
	default:
		return nil, status.Errorf(codes.NotFound, "not found type:%v", req.Name)
	}
	return resp, nil
}

func server() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServiceServer(s, &service{})
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func peer() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cc := echo.NewEchoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	parse := func(res *echo.EchoResp) {
		if res == nil {
			return
		}
		log.Println("resp desc:", res.Desc)

		if ptypes.Is(res.Character, &echo.Student{}) {
			student := &echo.Student{}
			if err := ptypes.UnmarshalAny(res.Character, student); err != nil {
				log.Fatal(err)
			}
			log.Printf("student info -> %+v \n", student)
		}

		if ptypes.Is(res.Character, &echo.Teacher{}) {
			teacher := &echo.Teacher{}
			if err := ptypes.UnmarshalAny(res.Character, teacher); err != nil {
				log.Fatal(err)
			}
			log.Printf("teacher info -> %+v \n", teacher)
		}
	}

	resp, err := cc.Echo(ctx, &echo.EchoReq{Name: "abc"})
	if err != nil {
		log.Printf("resp error: %s", err)
	}
	parse(resp)

	resp, err = cc.Echo(ctx, &echo.EchoReq{Name: "student"})
	if err != nil {
		log.Printf("resp error: %s", err)
	}
	parse(resp)

	resp, err = cc.Echo(ctx, &echo.EchoReq{Name: "teacher"})
	if err != nil {
		log.Printf("resp error: %s", err)
	}
	parse(resp)

}
