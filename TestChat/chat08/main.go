package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

func main() {
	test3()
}

func test1() {

	var elem atomic.Value
	type a struct {
		t uint32
		v string
	}

	if x, ok := elem.Load().(*a); ok {
		fmt.Println("x => ", x)
	} else {
		elem.Store(&a{123, "hello"})
	}

	if x, ok := elem.Load().(*a); ok {
		fmt.Println("x => ", x)
	} else {
		panic("test failed")
	}

}

func test2() {

	type a struct {
		sync.Mutex
		data map[int]string
	}

	var elem atomic.Value
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for k := 0; k < 1000; k++ {
				if x, ok := elem.Load().(*a); ok {
					x.Lock()
					x.data[rand.Int()] = "hello world"
					x.Unlock()
				} else {
					elem.Store(&a{data: make(map[int]string, 64)})
				}
			}
		}()
	}
	wg.Wait()

}

func test3() {

	type t struct {
		sync.Mutex
		data map[int]int
	}

	var elem atomic.Value
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for k := 0; k < 1000; k++ {

				if k%4 == 0 {
					tt := &t{
						data: make(map[int]int, 16),
					}
					tt.data[rand.Int()] = rand.Int()
					elem.Store(tt)
				}

				if x, ok := elem.Load().(*t); ok {
					_ = x.data[rand.Int()]
				}
			}
		}()
	}
	wg.Wait()

}
