package main

import (
	"log"
	"sync"
)

type self struct {
	data []string
}

func main() {

	var items = make([][]string, 0, 4)

	p := sync.Pool{New: func() interface{} {
		return &self{data: make([]string, 0, 3)}
	}}

	t1 := []string{"A", "B", "C"}
	a := p.Get().(*self)
	a.data = append(a.data, t1...)
	items = append(items, a.data)
	a.data = a.data[:0]
	p.Put(a)

	for k, v := range items {
		log.Printf("k %d v : %v ", k, v)
	}

	t2 := []string{"D", "E", "F"}
	b := p.Get().(*self)
	b.data = append(b.data, t2...)
	items = append(items, b.data)
	b.data = b.data[:0]
	p.Put(b)

	for k, v := range items {
		log.Printf("k %d v : %v ", k, v)
	}

	/*
		//a 和 b 用的是sync.Pool的同一块内存
		output:
		//第一次打印
		2021/06/02 22:17:58 k 0 v : [A B C]
		//第二次打印ABC变成DEF
		2021/06/02 22:17:58 k 0 v : [D E F]
		2021/06/02 22:17:58 k 1 v : [D E F]
	*/

}
