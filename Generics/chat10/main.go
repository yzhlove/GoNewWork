package main

import (
	"fmt"
	"reflect"
)

func main() {

	type req struct {
		a int
		b int
	}

	r := &req{
		a: 100,
		b: 50,
	}

	x := add[*req](r, func(r *req) *req {
		if r == nil {
			return nil
		}
		if r.a%r.b == 0 {
			return nil
		}
		return r
	})

	if x == nil {
		fmt.Println("ok ,x is nil")
	}

	fmt.Println(reflect.TypeOf(x), " -> ", x)
}

func add[T any](v T, tr func(v T) T) T {
	if tr != nil {
		return tr(v)
	}
	var c T
	return c
}
