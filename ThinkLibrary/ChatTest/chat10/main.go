package main

import (
	"fmt"
)

func main() {
	cc(100)
	fmt.Println()

	cc(200)
	fmt.Println()
	fmt.Println("ccccc")

	ee(1000)
}

func g(a int) bool {
	return a == 100
}

func aa(a int) {

	if g(a) {
		fmt.Println("return .")
		return
	}

	defer func() {
		fmt.Println("defer run:aaaa")
	}()

}

func cc(a int) {
	aa(a)
}

func ee(a int) {
	defer func() {
		fmt.Println("panic is before")
	}()
	panic("test defer call")
	defer func() {
		fmt.Println("panic is after")
	}()
}
