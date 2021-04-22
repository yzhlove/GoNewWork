package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := ret(); err != nil {
		panic(err)
	}
	fmt.Println("Ok.")
}

func retErr() (x int, err error) {
	x = 10
	err = errors.New("test output error")
	return
}

func ret() (err error) {
	if x, err := retErr(); err != nil {
		fmt.Println("x-> ", x)
	}
	return
}
