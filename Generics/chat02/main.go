package main

import (
	"fmt"
	"math/rand"
)

type data[K comparable, V int | string] map[K]V

func (d data[K, V]) Iterator() {
	fmt.Println("d len = ", len(d))
	for k, v := range d {
		fmt.Println(k, v)
	}
}

func (d data[K, V]) Callback(fn func(K, V)) {
	for k, v := range d {
		if fn != nil {
			fn(k, v)
		}
	}
}

func main() {

	var t1 data[string, int] = map[string]int{
		"a": 1,
		"b": 2,
	}

	fmt.Println(t1)

	var t2 data[string, int] = make(map[string]int, 15)
	for k, v := range rand.Perm(15) {
		t2[fmt.Sprintf("key-%d", k)] = v
	}

	fmt.Println(t2)
	t2.Iterator()

	t2.Callback(func(key string, value int) {
		fmt.Println("callback --> ", key, value)
	})

}
