package main

import (
	"errors"
	"fmt"
)

func main() {
	//if err := ret(); err != nil {
	//	panic(err)
	//}
	//fmt.Println("Ok.")
	retTest()
}

func retErr() (x int, err error) {
	x = 10
	err = errors.New("test output error")
	return
}

func ret() (err error) {
	//x , err := retErr()
	if x, err := retErr(); err != nil {
		fmt.Println("x-> ", x)
	}
	return
}

func retTest() {

	x, err := retErr()
	fmt.Println("0 x ->", &x, " err->", &err)
	if err != nil {
		x, err := retErr()
		fmt.Println("1 x ->", &x, " err->", &err)
	}
	fmt.Println("2 x ->", &x, " err->", &err)
	x, err = retErr()
	fmt.Println("3 x ->", &x, " err->", &err)
	y, err := retErr()
	fmt.Println("4 x ->", &y, " err->", &err)

}
