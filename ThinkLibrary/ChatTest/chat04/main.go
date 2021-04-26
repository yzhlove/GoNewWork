package main

import (
	"fmt"
	"sync"
	"time"
)

//尽量不要在锁里面使用syscall(系统调用)

var lock sync.RWMutex
var wg sync.WaitGroup

func main() {

	begin := time.Now()
	
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go gets()
	}
	wg.Wait()

	fmt.Println("over -> ", time.Since(begin))
}

func gets() {
	for i := 0; i < 100000; i++ {
		get(i)
	}
	wg.Done()
}

func get(i int) {
	beginTime := time.Now()
	lock.RLock()
	time.Since(beginTime).Nanoseconds()
	lock.RUnlock()
}
