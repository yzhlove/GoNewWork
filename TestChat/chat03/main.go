package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)

	start := time.Now()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("context done -> ", ctx.Err())
		case <-timer.C:
			fmt.Println("time stop -> ")
			cancel()
		}
		fmt.Println("run time -> ", time.Now().Sub(start).String())
	}()

	time.Sleep(time.Second)

}
