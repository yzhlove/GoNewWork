package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	sg := Group{}
	var count int32

	var succeed, failed, number int32

	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt32(&number, 1)
			_, err := sg.Do("name", func() (interface{}, error) {
				atomic.AddInt32(&count, 1)
				time.Sleep(time.Second)
				return "aaa", nil
			})
			if err != nil {
				atomic.AddInt32(&succeed, 1)
			} else {
				atomic.AddInt32(&failed, 1)
			}
		}()
	}

	fmt.Println(number, succeed, failed, count)
	//953 951 0 0
}

type Group struct {
	mutex sync.Mutex
	data  map[string]*call
}

type call struct {
	val  interface{}
	err  error
	stop chan struct{}
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error) {

	if g.mutex.TryLock() {
		defer g.mutex.Unlock()

		if g.data == nil {
			g.data = make(map[string]*call)
		}

		c := new(call)
		g.data[key] = c

		c.val, c.err = fn()
		return
	}
	return nil, errors.New("not take locker")
}
