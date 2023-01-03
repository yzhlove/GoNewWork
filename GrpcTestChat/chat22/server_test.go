package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test-chat/chat22/echo"
	"testing"
)

func Test_Client(t *testing.T) {

	cc, err := grpc.Dial("localhost:8887", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}

	conn := echo.NewEchoServiceClient(cc)
	if _, err := conn.Echo(context.Background(), &echo.Msg_Nil{}); err != nil {
		t.Error(err)
	} else {
		t.Log("Ok.")
	}
}

func Benchmark_Client1(b *testing.B) {

	cc, err := grpc.Dial("localhost:8887", grpc.WithInsecure())
	if err != nil {
		b.Error(err)
		return
	}

	conn := echo.NewEchoServiceClient(cc)

	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		if _, err := conn.Echo(context.Background(), &echo.Msg_Nil{}); err != nil {
			b.Error(err)
			return
		}
	}

	b.StopTimer()
	b.ReportAllocs()

}

func Benchmark_Client2(b *testing.B) {

	cc, err := grpc.Dial("localhost:8887", grpc.WithInsecure())
	if err != nil {
		b.Error(err)
		return
	}

	conn := echo.NewEchoServiceClient(cc)

	b.RunParallel(func(pb *testing.PB) {
		b.ResetTimer()
		b.StartTimer()
		for pb.Next() {
			if _, err := conn.Echo(context.Background(), &echo.Msg_Nil{}); err != nil {
				b.Error(err)
				return
			}
		}

		b.StopTimer()
		b.ReportAllocs()
	})

}
