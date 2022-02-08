package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	test2()
}

func test1() {
	var a uint32
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println("value before => ", atomic.LoadUint32(&a))
			atomic.AddUint32(&a, 1)
			fmt.Println("value after => ", atomic.LoadUint32(&a))
			fmt.Println()
		}()
	}
	wg.Wait()

	fmt.Println("value => ", a)

}

func test2() {

	var a uint32
	var wg sync.WaitGroup

	numCh := make(chan uint32, 100)
	go func() {
		rept := make(map[uint32]struct{}, 100)
		for num := range numCh {
			if _, ok := rept[num]; ok {
				fmt.Println("rept value ==> ", num)
			} else {
				rept[num] = struct{}{}
			}
		}
	}()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			value := atomic.LoadUint32(&a)
			atomic.CompareAndSwapUint32(&a, value, value+1)
			numCh <- atomic.LoadUint32(&a)
		}()
	}
	wg.Wait()

	fmt.Println("\n\n a value ==> ", a)
}
