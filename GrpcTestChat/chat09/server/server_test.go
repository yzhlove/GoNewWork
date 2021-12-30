package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-test-chat/chat09/echo"
	"sync"
	"testing"
)

func Test_Client(t *testing.T) {

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cc, err := grpc.Dial(fmt.Sprintf("localhost%s", _port), grpc.WithInsecure())
			if err != nil {
				t.Error(err)
				return
			}
			conn := echo.NewEchoServiceClient(cc)
			resp, err := conn.Echo(context.Background(), &echo.Req{Msg: fmt.Sprintf("test-%d", i)})
			if err != nil {
				t.Error(err)
				return
			}
			t.Log("resp:", resp.Msg)
		}(i + 1)
	}
	wg.Wait()

}
