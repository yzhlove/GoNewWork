package main

import (
	"context"
	"fmt"
	"hash/crc32"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	_MaxQueue = runtime.NumCPU() * 2
	_MaxChan  = 128
)

type Reqer interface {
	GetNumber() int
	GetMemory() any
}

type Req struct {
	str1, str2 string
}

func (r *Req) GetNumber() int {
	return int(crc32.ChecksumIEEE([]byte(r.str1 + r.str2)))
}

func (r *Req) GetMemory() any {
	return &Req{}
}

type Resp struct {
	number uint32
	err    error
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	manager := NewManager[*Req, *Resp](ctx, func(req *Req) *Resp {
		fmt.Println("req => ", req.str1, req.str2)
		var resp = &Resp{}
		var1, err := strconv.ParseUint(req.str1, 10, 32)
		if err != nil {
			resp.err = err
			return resp
		}
		var2, err := strconv.ParseUint(req.str2, 10, 32)
		if err != nil {
			resp.err = err
			return resp
		}
		resp.number = uint32(var1 + var2)
		return resp
	})

	if err := manager.run(); err != nil {
		return
	}

	var cases = []struct {
		a string
		b string
	}{
		{"1", "2"},
		{"3", "4"},
		{"5", "6"},
		{"7", "8"},
		{"9", "10"},
		{"11", "12"},
		{"13", "14"},
		{"a12", "a14"},
		{"x13", "ff14"},
		{"t13", "ggg14"},
	}

	var wg sync.WaitGroup
	for k := range cases {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			resp := make(chan *Resp)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			x := manager.Get()
			x.str1, x.str2 = cases[k].a, cases[k].b

			manager.Do(x, resp)
			select {
			case <-ctx.Done():
				fmt.Println("timeout k ==> ", k)
			case r := <-resp:
				fmt.Println("result ==> ", r.number, r.err)
			}
		}(k)
	}
	wg.Wait()
	cancel()
}

type reqPack[T Reqer, V any] struct {
	req T
	ch  chan V
}

type Manager[T Reqer, V any] struct {
	ctx context.Context
	sync.Pool
	eventQueue []chan *reqPack[T, V]
	eventFunc  func(T) V
}

func NewManager[T Reqer, V any](ctx context.Context, eventFunc func(T) V) *Manager[T, V] {
	manager := &Manager[T, V]{ctx: ctx, eventFunc: eventFunc}
	manager.eventQueue = make([]chan *reqPack[T, V], _MaxQueue)
	manager.New = func() any {
		var ct T
		return ct.GetMemory()
	}
	return manager
}

func (manager *Manager[T, V]) Do(req T, ch chan V) {
	k := req.GetNumber() % _MaxQueue
	manager.eventQueue[k] <- &reqPack[T, V]{req: req, ch: ch}
}

func (manager *Manager[T, V]) Get() T {
	return manager.Pool.Get().(T)
}

func (manager *Manager[T, V]) run() error {
	for k := range manager.eventQueue {
		manager.eventQueue[k] = make(chan *reqPack[T, V], _MaxChan)
		go func(k int) {
			for {
				select {
				case <-manager.ctx.Done():
					return
				case r := <-manager.eventQueue[k]:
					if manager.eventFunc != nil {
						r.ch <- manager.eventFunc(r.req)
						manager.Put(r.req)
					}
				}
			}
		}(k)
	}
	return nil
}
