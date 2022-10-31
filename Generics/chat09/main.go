package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	ta := NewA(123)
	fmt.Println(ta.count.Load())

	tb := *ta
	tb.count.Store(456)
	fmt.Println(tb.count.Load())
	fmt.Println(ta.count.Load())

}

type A struct {
	count atomic.Int32
}

func NewA(v int32) *A {
	a := &A{}
	a.count.Store(v)
	return a
}
