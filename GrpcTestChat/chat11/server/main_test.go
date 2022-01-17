package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-test-chat/chat11/echo"
	"math/rand"
	"strings"
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

func character(number int) string {
	arr := make([]string, 0, number)
	for i := 0; i < number; i++ {
		arr = append(arr, fmt.Sprintf("%d", rand.Intn(100)+10))
	}
	return strings.Join(arr, "-")
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
				&echo.Req{Msg: character(rand.Intn(50) + 20)})
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

// Benchmark_Handle-10    	    6016	    216953 ns/op	  122327 B/op	     410 allocs/op
// serial Benchmark_Handle-10    	    5947	    210935 ns/op	  122245 B/op	     411 allocs/op
// rand Benchmark_Handle-10    	    5372	    224998 ns/op	  124143 B/op	     458 allocs/op
// rand Benchmark_Handle-10    	    5923	    214063 ns/op	  123553 B/op	     459 allocs/op
