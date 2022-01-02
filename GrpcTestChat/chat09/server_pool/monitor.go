package main

import (
	"strconv"
	"sync"
)

type monitor struct {
	collection []string
}

func (m *monitor) Push(strs ...string) {
	m.collection = append(m.collection, strs...)
}

func (m *monitor) Merge() string {
	if len(m.collection) > 0 {
		var sum int64
		for _, v := range m.collection {
			number, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				continue
			}
			sum += number
		}
		return strconv.FormatInt(sum, 10)
	}
	return "0"
}

func (m *monitor) Clear() {
	if len(m.collection) > 0 {
		m.collection = m.collection[:0]
	}
}

type poolMonitor struct {
	pool *sync.Pool
}

func New() *poolMonitor {
	return &poolMonitor{
		pool: &sync.Pool{New: func() interface{} {
			return &monitor{collection: make([]string, 0, 4)}
		}},
	}
}

func (p *poolMonitor) Get() *monitor {
	res := p.pool.Get()
	return res.(*monitor)
}

func (p *poolMonitor) Put(m *monitor) {
	if len(m.collection) > 0 {
		m.Clear()
		p.pool.Put(m)
	}
}
