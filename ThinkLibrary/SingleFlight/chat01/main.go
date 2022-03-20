package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	sg := Group{}
	var count int32

	for i := 0; i < 100; i++ {
		go func() {
			sg.Do("name", func() (interface{}, error) {
				atomic.AddInt32(&count, 1)
				time.Sleep(time.Second)
				return "hello world", nil
			})
		}()
	}

	if atomic.LoadInt32(&count) != 1 {
		panic(any("count invalid"))
	} else {
		fmt.Println("ok")
	}

}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error) {
	g.mu.Lock()

	if g.m == nil {
		g.m = make(map[string]*call)
	}

	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()

		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err

}
