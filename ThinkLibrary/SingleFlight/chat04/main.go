package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	sg := Group{data: make(map[string]*call)}
	var count int32

	for i := 0; i < 1000; i++ {
		go func() {
			sg.Do("name", func() (interface{}, error) {
				atomic.AddInt32(&count, 1)
				time.Sleep(time.Second)
				return "hello world", nil
			})
		}()
	}

	fmt.Println("count => ", count)

}

type Group struct {
	mutex sync.Mutex
	data  map[string]*call
}

type call struct {
	stop chan struct{}
	val  interface{}
	err  error
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mutex.Lock()

	if c, ok := g.data[key]; ok {
		g.mutex.Unlock()
		<-c.stop
		return c.val, c.err
	}

	c := new(call)
	c.stop = make(chan struct{})
	g.data[key] = c
	g.mutex.Unlock()

	c.val, c.val = fn()
	close(c.stop)

	g.mutex.Lock()
	delete(g.data, key)
	g.mutex.Unlock()
	return c.val, c.err
}
