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

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	manager := NewManager(ctx, func(req *Req) (resp Resp) {
		var1, err := strconv.ParseUint(req.str1, 10, 32)
		if err != nil {
			resp.err = err
			return
		}
		var2, err := strconv.ParseUint(req.str2, 10, 32)
		if err != nil {
			resp.err = err
			return
		}
		resp.number = uint32(var1 + var2)
		return
	})
	manager.run()

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
			resp := make(chan Resp)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			manager.Do(&Req{
				str1: cases[k].a,
				str2: cases[k].b,
			}, resp)

			select {
			case <-ctx.Done():
				fmt.Println("k timeout", k)
			case r := <-resp:
				fmt.Println("result -> ", r.number, r.err)
			}
		}(k)
	}

	wg.Wait()
	cancel()
}

type Req struct {
	str1, str2 string
}

type Pack struct {
	resp chan Resp
	req  *Req
}

type Resp struct {
	number uint32
	err    error
}

var (
	_MaxQueue = runtime.NumCPU() * 2
	_MaxChan  = 128
)

type Manager struct {
	ctx context.Context
	sync.Pool
	eventQueue []chan *Pack
	eventFunc  func(req *Req) Resp
}

func (manager *Manager) Do(req *Req, ch chan Resp) {
	index := int(crc32.ChecksumIEEE([]byte(req.str1+req.str2))) % _MaxQueue
	manager.eventQueue[index] <- &Pack{req: req, resp: ch}
}

func (manager *Manager) run() error {
	for k := range manager.eventQueue {
		manager.eventQueue[k] = make(chan *Pack, _MaxChan)
		go func(k int) {
			for {
				select {
				case <-manager.ctx.Done():
					return
				case pack := <-manager.eventQueue[k]:
					if manager.eventFunc != nil {
						pack.resp <- manager.eventFunc(pack.req)
						manager.Put(pack.req)
					}
				}
			}
		}(k)
	}
	return nil
}

func NewManager(ctx context.Context, eventFunc func(req *Req) Resp) *Manager {
	manager := &Manager{ctx: ctx, eventFunc: eventFunc}
	manager.eventQueue = make([]chan *Pack, _MaxQueue)
	manager.New = func() any {
		return &Req{}
	}
	return manager
}
