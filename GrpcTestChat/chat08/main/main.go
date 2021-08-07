package main

import (
	"context"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-test-chat/chat08/echo"
)

const (
	stuName   = "student"
	teachName = "teacher"
)

func main() {

}

type service struct{}

func (s *service) Echo(ctx context.Context, req *echo.Req) (*echo.Resp, error) {
	resp := &echo.Resp{Desc: "service:" + req.Type}
	switch req.Type {
	case stuName:
		student := &echo.Student{Name: "大宝", Subject: "数学"}
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
	return nil, nil
}
