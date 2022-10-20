package main

import (
	"context"
	"hash/crc32"
	"sync"
	"testing"
)

func Benchmark_Manager(b *testing.B) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	manager := NewManager(ctx, func(req *Req) (resp Resp) {
		resp.number = crc32.ChecksumIEEE([]byte(req.str1 + req.str2))
		return
	})
	manager.run()

	b.StartTimer()

	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp := make(chan Resp)
			manager.Do(&Req{
				str1: "a",
				str2: "b",
			}, resp)
			<-resp
		}()
	}
	wg.Wait()
	b.StartTimer()
	b.ReportAllocs()
}
