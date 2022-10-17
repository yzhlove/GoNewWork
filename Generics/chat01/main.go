package main

import "fmt"

func main() {

	type slice[T int | string] []T

	ss := make(slice[string], 0, 3)
	ss = append(ss, "hello")
	ss = append(ss, "world")

	fmt.Println(ss)

	ss2 := make(slice[int], 0, 3)

	for i := 0; i < 4; i++ {
		ss2 = append(ss2, i)
	}

	fmt.Println(ss2)

}
