package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/test/bufconn"
	"grpc-test-chat/chat06/contions/echo"
	"log"
	"net"
	"testing"
	"time"
)

var listener *bufconn.Listener

func initGppcServer() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServiceServer(s, &service{})
	reflection.Register(s)
	go func() {
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()
}

func getBufDialer(listener *bufconn.Listener) func(ctx context.Context, str string) (net.Conn, error) {
	return func(ctx context.Context, str string) (net.Conn, error) {
		return listener.Dial()
	}
}

func initGrpcServerBuffer() {
	listener = bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	echo.RegisterEchoServiceServer(s, &service{})
	reflection.Register(s)
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
}

func Test_Echo(t *testing.T) {
	initGppcServer()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	cc := echo.NewEchoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := cc.Echo(ctx, &echo.EchoReq{Name: "student"})
	if err != nil {
		t.Error(err)
		return
	}
	student := &echo.Student{}
	if ptypes.Is(resp.Character, student) {
		if err := ptypes.UnmarshalAny(resp.Character, student); err != nil {
			t.Error(err)
		}
	}
	t.Log("resp ", resp.Desc, resp.Character.TypeUrl, student.Name, student.Class)
}

func Test_EchoBuff(t *testing.T) {
	initGrpcServerBuffer()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "bufNet", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}

	newCtx, newCancel := context.WithTimeout(context.Background(), time.Second)
	defer newCancel()

	cc := echo.NewEchoServiceClient(conn)
	resp, err := cc.Echo(newCtx, &echo.EchoReq{Name: "teacher"})
	if err != nil {
		t.Error(err)
		return
	}

	teacher := &echo.Teacher{}
	if ptypes.Is(resp.Character, teacher) {
		if err := ptypes.UnmarshalAny(resp.Character, teacher); err != nil {
			t.Error(err)
			return
		}
	}
	t.Log(resp.Desc, resp.Character.TypeUrl, teacher.Name, teacher.Project, teacher.Salary)
}
