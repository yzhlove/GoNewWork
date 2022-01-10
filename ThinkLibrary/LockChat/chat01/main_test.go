package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Test_SpinLocker(t *testing.T) {

	locker := NewSpinLocker()

	var count = 10000
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			locker.Lock()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(50)+10))
			locker.Unlock()

		}()
	}
	wg.Wait()

	t.Log("Ok")
}

func Benchmark_SpinLocker(b *testing.B) {

	locker := NewSpinLocker()

	b.StartTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

			locker.Lock()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(50)+10))
			locker.Unlock()

		}
	})

	b.StopTimer()
	b.ReportAllocs()

}

// Benchmark_SpinLocker-10    	     344	   3489236 ns/op	      17 B/op	       0 allocs/op time
// Benchmark_SpinLocker-10    	 6345738	       190.6 ns/op	       0 B/op	       0 allocs/op no_time
