package main

import (
	"math/rand"
	"testing"
	"time"
)

func Benchmark_Sync(b *testing.B) {

	rand.Seed(time.Now().UnixNano())
	st := New()

	str := "test"

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		st.set(str, rand.Perm(rand.Intn(100))...)
		st.add(str, rand.Intn(10000))
	}

	b.StopTimer()
	b.ReportAllocs()

}
