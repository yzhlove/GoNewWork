package main

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"grpc-test-chat/chat08/echo"
	"testing"
	"time"
)

type rpcData struct {
	data proto.Message
}

func (r *rpcData) Matches(msg interface{}) bool {
	res, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(res, r.data)
}

func (r *rpcData) String() string {
	return fmt.Sprintf("is %s", r.data)
}

func Test_Client(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	echocc := NewMockHelloServiceClient(ctrl)

	req := &echo.Req{Type: stuName}
	any, err := types.MarshalAny(&echo.Student{Name: "大宝", Subject: "数学", Address: "China Shanghai"})
	if err != nil {
		t.Error(err)
		return
	}
	resp := &echo.Resp{Desc: "mock result", Character: any}
	echocc.EXPECT().Echo(gomock.Any(), &rpcData{data: req}).Do(func(ctx context.Context, in interface{}, opts ...interface{}) {
		if res, ok := in.(proto.Message); ok {
			fmt.Println("proto.Message => ", res)
		} else {
			fmt.Println("type assert error:", in)
		}
	}).Return(resp, nil).AnyTimes()

	testEcho(t, echocc)

}

func testEcho(t *testing.T, mock *MockHelloServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := mock.Echo(ctx, &echo.Req{Type: stuName})
	if err != nil {
		t.Fatal("mock call error:", err)
	}

	if resp.Desc != "mock result" {
		t.Error("mock rest error")
		return
	}

	if types.Is(resp.Character, &echo.Student{}) {
		student := &echo.Student{}
		if err := types.UnmarshalAny(resp.Character, student); err != nil {
			t.Error(err)
		}
		t.Log("mock test succeed:", student.Name, student.Subject, student.Address)
	}

}
