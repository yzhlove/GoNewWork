package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	contextWithDeadline2()
}

func contextWithDeadline() {

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("stop ... ")
		cancel()
	}()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			nctx, can := context.WithTimeout(ctx, time.Duration(i+1)*time.Second)
			fmt.Printf("start [%d] \n", i)
			<-nctx.Done()
			fmt.Printf("stop [%d] \n", i)
			can()
		}(i)
	}
	wg.Wait()
}

func contextWithDeadline2() {

	var ctx = context.Background()
	var cancel context.CancelFunc
	var st = time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		ctx, cancel = context.WithTimeout(ctx, time.Duration((i+1)*10)*time.Second)
		go func(ctx context.Context, cancel context.CancelFunc, i int) {
			defer wg.Done()
			fmt.Println("start: ", i)
			<-ctx.Done()
			fmt.Println("over: ", i, time.Now().Sub(st).String())
			if i == 5 {
				cancel()
				fmt.Println("cancel ...")
			}
		}(ctx, cancel, i)
	}
	wg.Wait()
}
