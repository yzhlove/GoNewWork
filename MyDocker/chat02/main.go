package main

import (
	"fmt"
	"my-docker/chat02/a"
	"my-docker/chat02/b"
)

func main() {

	aa := &a.A{}
	bb := &b.B{}

	aa.Init(bb)
	bb.Init(aa)

	fmt.Println(bb.Obt().Run())

}
