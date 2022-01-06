package main

import (
	"fmt"
	"sync"
)

type poolMutex struct {
	pool *sync.Pool
}

func New() *poolMutex {
	return &poolMutex{
		pool: &sync.Pool{
			New: func() interface{} {
				return &sync.Mutex{}
			},
		},
	}
}

func (p *poolMutex) Get() *sync.Mutex {
	return p.pool.Get().(*sync.Mutex)
}

func (p *poolMutex) Put(m *sync.Mutex) {
	p.pool.Put(m)
}

func main() {
	fmt.Println("hello world")
}
