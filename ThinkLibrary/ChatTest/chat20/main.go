package main

import "fmt"

func main() {

	//var _ A = (*base)(nil)
	var _ A = (*B)(nil)

}

type A interface {
	sub()
	add()
	ok()
	enough()
	have()
}

type base struct{}

func (base) sub() {}

func (base) add() {}

func (base) enough() {}

func (base) have() {}

type B struct {
	base
}

func (b B) ok() {
	fmt.Println("a")
}
