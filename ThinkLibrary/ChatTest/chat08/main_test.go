package main

import (
	"math/rand"
	"testing"
)

func Test_Get(t *testing.T) {

	n := 1000
	s := new(n)

	for i := 0; i < n; i++ {
		s.put(rand.Float32())
	}

	for i := 0; i < n; i++ {
		t.Log(s.GetInt(1, 100000))
	}

}

func Benchmark_Get(b *testing.B) {

	s := new(1000)
	stat := make(chan struct{})
	go func() {
		for {
			select {
			case <-stat:
				return
			default:
				s.put(rand.Float32())
			}
		}
	}()

	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s.GetInt(1, 10000)
	}
	b.StopTimer()
	close(stat)
	b.ReportAllocs()
}

//-spin Benchmark_Get-12    	 5670387	       208.2 ns/op	      19 B/op	       0 allocs/op
//-rw   Benchmark_Get-12    	 2571800	       452.9 ns/op	      21 B/op	       0 allocs/op
