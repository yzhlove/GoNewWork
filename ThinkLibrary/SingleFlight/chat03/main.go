package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	var count int32
	//var mutex sync.Mutex
	var stop chan struct{}

	stop = make(chan struct{})

	var s, f, n int32

	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt32(&n, 1)
			//if mutex.TryLock() {
			//	atomic.AddInt32(&s, 1)
			//	defer mutex.Unlock()
			//	atomic.AddInt32(&count, 1)
			//	time.Sleep(time.Second)
			//	close(stop)
			//} else {
			//	atomic.AddInt32(&f, 1)
			//}
		}()
	}

	<-stop
	fmt.Println(s, f, n, count)
}
