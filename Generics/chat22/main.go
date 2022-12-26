package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var wg sync.WaitGroup

	var count atomic.Int32

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Add(1)
			if count.Load() > 25 {
				defer count.Add(-1)
			}
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		}()
	}

	wg.Wait()
	fmt.Println("count --> ", count.Load())

}
