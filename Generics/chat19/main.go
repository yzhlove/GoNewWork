package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go func() {
		var index = 0
		tick := time.NewTicker(time.Second)
		for {
			<-tick.C
			index++
			ch <- index
			if index >= 5 {
				close(ch)
				break
			}
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-ch:
			fmt.Println("get chan ", v, ok)
		}
	}

}
