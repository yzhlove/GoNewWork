package main

import (
	"fmt"
	"sync"
)

func main() {

	x := pool.Get().(*req)
	y := pool.Get().(*req)
	pool.Put(x)
	z := pool.Get().(*req)
	y = y
	z = z
}

type req struct {
	a, b string
}

var pool = sync.Pool{
	New: func() any {
		fmt.Println("---------- pool.new")
		return &req{}
	},
}
