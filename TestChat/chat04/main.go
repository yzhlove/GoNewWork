package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	msgch := make(chan string)
	go func() {
		for msg := range msgch {
			fmt.Println("msg --> ", msg)
		}
	}()
	go func() {
		fmt.Println("1 status -> ", sendMsgWithContext(context.Background(), time.Second, "hello", msgch))
	}()
	go func() {
		fmt.Println("2 status -> ", sendMsgWithTimer(time.Second, "world", msgch))
	}()

	time.Sleep(time.Second)
	
}

func sendMsgWithContext(ctx context.Context, t time.Duration, msg string, ch chan<- string) bool {
	ctx, cancel := context.WithTimeout(ctx, t)
	defer cancel()
	select {
	case <-ctx.Done():
		return false
	case ch <- msg:
		return true
	}
}

func sendMsgWithTimer(t time.Duration, msg string, ch chan<- string) bool {
	tm := time.NewTimer(t)
	select {
	case <-tm.C:
		tm.Stop()
		return false
	case ch <- msg:
		return true
	}
}
