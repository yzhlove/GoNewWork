package main

import "sync"

func main() {

}

type mp struct {
	*sync.Pool
}

func New(size int) *mp {
	if size == 0 {
		size = 16
	}
	return &mp{
		&sync.Pool{
			New: func() interface{} {
				return make(map[string]string, size)
			},
		},
	}
}

func (m *mp) Get() map[string]string {
	return m.Pool.Get().(map[string]string)
}

func (m *mp) Put(data map[string]string) {
	for k := range data {
		delete(data, k)
	}
	m.Pool.Put(data)
}
