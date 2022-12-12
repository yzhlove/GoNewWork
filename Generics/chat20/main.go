package main

import (
	"fmt"
	"sort"
)

func main() {

	s1 := inMemSort(arraySource(1, 3, 1, 4, 5, 2, 1, 7, 6, 3, 8, 4))
	s2 := inMemSort(arraySource(4, 4, 3, 5, 2, 1, 7, 6, 3, 7))
	for x := range merge(s1, s2) {
		fmt.Print(x, ",")
	}
}

func arraySource(values ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range values {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func inMemSort(ch <-chan int) <-chan int {
	sch := make(chan int)
	go func() {
		values := make([]int, 0, 64)
		for v := range ch {
			values = append(values, v)
		}
		sort.Ints(values)
		for _, x := range values {
			sch <- x
		}
		close(sch)
	}()
	return sch
}

func merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}
