package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	test()
}

func test() {

	var a uint32
	var cnt uint32
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		var bbb sync.WaitGroup
		for i := 0; i < 1000; i++ {
			bbb.Add(1)
			go func() {
				defer bbb.Done()
				if value := a; value > 500 {
					atomic.AddUint32(&cnt, 1)
				}
			}()
		}
		bbb.Wait()
	}()

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

	wg.Wait()
	fmt.Println("a => ", a, " cnt => ", cnt)
}
