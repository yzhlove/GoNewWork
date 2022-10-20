package main

import (
	"fmt"
	"sync"
)

func main() {

	x := pool.Get()
	if x == nil {
		value := []byte("hello world")
		pool.Put(value)
	}

	x = pool.Get()
	if x != nil {
		ret, ok := x.([]byte)
		if ok {
			fmt.Println(string(ret))
		}
	} else {
		fmt.Println("error")
	}
}

var pool = sync.Pool{}
