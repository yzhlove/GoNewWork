package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 关于atomic的实验

func main() {

	var a uint32
	var wg sync.WaitGroup

	reptCh := make(chan uint32, 100)
	go func() {
		reptValue := make(map[uint32]struct{}, 100)
		for value := range reptCh {
			if _, ok := reptValue[value]; ok {
				fmt.Println("rept value => ", value)
			} else {
				reptValue[value] = struct{}{}
			}
		}
	}()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint32(&a, 1)
			reptCh <- atomic.LoadUint32(&a)

		}()
	}
	wg.Wait()
	close(reptCh)

	fmt.Println("a value => ", a)
}
