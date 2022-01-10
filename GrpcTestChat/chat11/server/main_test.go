package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test-chat/chat11/echo"
	"sync"
	"testing"
)

func Test_Handle(t *testing.T) {
	cc, err := grpc.Dial("localhost:55566", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}

	conn := echo.NewEchoServiceClient(cc)

	resp, err := conn.Echo(context.Background(), &echo.Req{Msg: "1-2-3-4-5-6-7-8-9-10"})
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("[%d] resp => %s \n", resp.Code, resp.Msg)
}

func Test_Handle2(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			cc, conn, err := taskConn()
			if err != nil {
				t.Error(err)
				return
			}

			resp, err := conn.Echo(context.Background(), &echo.Req{Msg: "1-2-3"})
			if err != nil {
				t.Error(err)
				return
			}

			t.Logf("resp =>  [%s] [%d]", resp.Msg, resp.Code)

			cc.Close()

		}()
	}
	wg.Wait()

}

func taskConn() (*grpc.ClientConn, echo.EchoServiceClient, error) {
	cc, err := grpc.Dial("localhost:55566", grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	conn := echo.NewEchoServiceClient(cc)
	return cc, conn, nil
}

func Benchmark_Handle(b *testing.B) {

	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {

		for pb.Next() {
			cc, conn, err := taskConn()
			if err != nil {
				b.Error(err)
				return
			}

			_, err = conn.Echo(context.Background(),
				&echo.Req{Msg: "1-2-3-4-5-6-7-8-9-10"})
			if err != nil {
				b.Error(err)
				return
			} else {

			}

			cc.Close()
		}

	})
	b.StopTimer()
	b.ReportAllocs()

}
