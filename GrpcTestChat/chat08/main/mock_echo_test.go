package main

import (
	"github.com/golang/mock/gomock"
	"grpc-test-chat/chat08/echo"
	"testing"
)

func Test_Client(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := NewMockHelloServiceClient(ctrl)
	req := &echo.Req{Type: stuName}

	client.EXPECT().Echo(gomock.Any(), req).Return(req)
}
