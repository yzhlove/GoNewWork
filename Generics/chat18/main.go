package main

import (
	"fmt"
	"sort"
)

func main() {

	for x := range inMemSort(arraySource(4, 3, 5, 7, 6, 2, 1, 8, 9, 0)) {
		fmt.Print(x)
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
