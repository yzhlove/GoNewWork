package main

import "log"

func main() {

	var a, b, c = 1, 2, 3
	a = b + c
	b = c + a
	c = a + b
	log.Println("a = ", a)
	log.Println("b = ", b)
	log.Println("c = ", c)
}
