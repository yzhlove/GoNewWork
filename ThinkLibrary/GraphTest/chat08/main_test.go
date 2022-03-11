package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func Test_MapPool(t *testing.T) {

	bucket := 500
	type kvData struct {
		k, v string
	}

	channel := make(chan *kvData, 1)
	stop := make(chan struct{})

	go func() {

		stat := make(map[string]int, bucket)
		for value := range channel {
			stat[value.v]++
		}
		for k, v := range stat {
			if v != 1000 {
				t.Errorf("error:k ===> %s ,number %d", k, v)
			}
		}
		close(stop)
	}()

	var wg sync.WaitGroup
	for i := 0; i < bucket; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			p := New(0)
			data := p.Get()
			value := fmt.Sprintf("%d", i)
			for k := 0; k < 1000; k++ {
				number := rand.Intn((i+1)*100) + i*100
				key := fmt.Sprintf("%d", number)
				data[key] = value
				channel <- &kvData{k: key, v: value}
			}
			p.Put(data)
		}(i)
	}
	wg.Wait()
	close(channel)

	<-stop
	fmt.Println("Ok.")
}

func Benchmark_Map(b *testing.B) {

	b.Run("map", func(b *testing.B) {
		b.ResetTimer()
		b.StartTimer()
		var count = 128
		for i := 0; i < b.N; i++ {
			data := make(map[string]string, count)
			for k := 0; k < count*10; k++ {
				data[strconv.Itoa(k)] = strconv.Itoa(k*10 + 1)
			}
		}
		b.StopTimer()
		b.ReportAllocs()
	})

	b.Run("map pool", func(b *testing.B) {
		b.ResetTimer()
		b.StartTimer()
		var count = 128
		p := New(count)
		for i := 0; i < b.N; i++ {
			data := p.Get()
			for k := 0; k < count*10; k++ {
				data[strconv.Itoa(k)] = strconv.Itoa(k*10 + 1)
			}
			p.Put(data)
		}
		b.StopTimer()
		b.ReportAllocs()
	})

	//Benchmark_Map/map-12         	    8704	    151240 ns/op	  168188 B/op	    2471 allocs/op
	//Benchmark_Map/map_pool
	//Benchmark_Map/map_pool-12    	   14470	     81180 ns/op	   11983 B/op	    2450 allocs/op

}

func Benchmark_MapParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var count = 128
		for pb.Next() {
			data := make(map[string]string, count)
			for k := 0; k < count*10; k++ {
				data[strconv.Itoa(k)] = strconv.Itoa(k*10 + 1)
			}
		}
	})
	b.ReportAllocs()
}

// Benchmark_MapParallel-12    	   23961	     50041 ns/op	  168189 B/op	    2471 allocs/op

func Benchmark_MapPoolParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var count = 128
		p := New(count)
		for pb.Next() {
			data := p.Get()
			for k := 0; k < count*10; k++ {
				data[strconv.Itoa(k)] = strconv.Itoa(k*10 + 1)
			}
			p.Put(data)
		}
	})
	b.ReportAllocs()
}

// Benchmark_MapPoolParallel-12    	   64810	     19761 ns/op	   12062 B/op	    2450 allocs/op
