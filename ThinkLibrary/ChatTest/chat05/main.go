package main

import (
	"fmt"
	"sync"
	"time"
)

var mlock sync.RWMutex
var mwg sync.WaitGroup

func main() {

	begin := time.Now()

	mwg.Add(100)
	for i := 0; i < 100; i++ {
		go gets()
	}
	mwg.Wait()

	fmt.Println("over:", time.Since(begin))
}

func gets() {
	defer mwg.Done()
	t := time.Now()
	for i := 0; i < 100000; i++ {
		get(i, &t)
	}
}

func get(i int, t *time.Time) {
	mlock.RLock()
	defer mlock.RUnlock()
	time.Since(*t).Nanoseconds()
}
