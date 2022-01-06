package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Benchmark_Mutex(b *testing.B) {

	pm := New()

	dosomething := func() {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+100))
	}

	b.StartTimer()

	b.RunParallel(func(pb *testing.PB) {

		for pb.Next() {

			mutex := pm.Get()
			mutex.Lock()
			dosomething()
			mutex.Unlock()
			pm.Put(mutex)

		}

	})

	b.StopTimer()
	b.ReportAllocs()

}

// Benchmark_Mutex-10    	      22	  51548309 ns/op	     161 B/op	       2 allocs/op

func Benchmark_New(b *testing.B) {
	pm := New()

	dosomething := func() {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+100))
	}

	b.StartTimer()

	var wg sync.WaitGroup

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mutex := pm.Get()
			mutex.Lock()
			dosomething()
			mutex.Unlock()
			pm.Put(mutex)

		}()
	}
	wg.Wait()
	b.StopTimer()
	b.ReportAllocs()
}

// Benchmark_New-10    	  249928	      4162 ns/op	     285 B/op	       3 allocs/op
