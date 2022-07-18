package main

import "fmt"

func main() {

	var a, b, c = 1, 2, 3
	a = b + c
	b = a + c
	c = a + b

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
}
