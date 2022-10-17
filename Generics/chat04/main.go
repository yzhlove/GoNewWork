package main

import (
	"fmt"
	"math/rand"
)

func main() {

	v, ok := Sp(128)
	fmt.Println(v, ok)

}

type value struct {
	a, b int
}

func Ret[T any](v T) (T, bool) {
	return v, true
}

func Sp(number int) (*value, bool) {

	x := &value{
		a: number,
		b: number - rand.Intn(number),
	}

	return Ret[*value](x)
}
