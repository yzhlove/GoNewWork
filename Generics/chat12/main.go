package main

import "fmt"

func main() {
	add(1, 2)
}

func add(a, b int) (c int) {
	defer func() {
		fmt.Println(1, " c = ", c)
	}()
	c = 10
	defer func() {
		fmt.Println(2, " c = ", c)
	}()
	return a + b
}
