package main

import (
	"fmt"
	"sync"
)

func main() {

	value := pool.Get().([]byte)
	copy(value, "hello")
	str := string(value)
	fmt.Printf("value:%p str:%p \n", value, &str)
	pool.Put(value)
	fmt.Printf("value:%s str:%s \n", string(value), str)

	value2 := pool.Get().([]byte)
	copy(value2, "world")
	fmt.Printf("value:%s str:%s \n", string(value), str)
	pool.Put(value2)

}

var pool = sync.Pool{
	New: func() any {
		return make([]byte, 8)
	},
}
