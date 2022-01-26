package main

import (
	"context"
	"grpc-test-chat/chat14/echo"
	"grpc-test-chat/chat14/grpcpool"
	"testing"
)

const (
	Host = "127.0.0.1:55566"
)

func Test_Echo(t *testing.T) {

	pool := grpcpool.New(Host, 4, 16)

	cc, err := pool.Get()
	if err != nil {
		t.Error(err)
		return
	}
	defer pool.Put(cc)

	client := echo.NewEchoServiceClient(cc.Conn)

	resp, err := client.Echo(context.Background(), &echo.Req{Msg: "1-2-3-4-5"})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("resp => ", resp.Msg)

}
