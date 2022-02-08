package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	test2()
}

func test() {

	var a uint32
	var cnt uint32
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		var aaa sync.WaitGroup

		for i := 0; i < 1000; i++ {
			aaa.Add(1)
			go func() {
				defer aaa.Done()
				atomic.AddUint32(&a, 1)
			}()
		}
		aaa.Wait()
	}()

	go func() {
		defer wg.Done()

		var bbb sync.WaitGroup
		for i := 0; i < 1000; i++ {
			bbb.Add(1)
			go func() {
				defer bbb.Done()
				if value := atomic.LoadUint32(&a); value > 500 {
					atomic.AddUint32(&cnt, 1)
				}
			}()
		}
		bbb.Wait()
	}()

	wg.Wait()
	fmt.Println("a => ", a, " cnt => ", cnt)
}

func test2() {
	var a uint32
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if value := atomic.LoadUint32(&a); value >= 500 {
				return
			}
			atomic.AddUint32(&a, 1)
		}()
	}
	wg.Wait()

	fmt.Println("value => ", a)
}
