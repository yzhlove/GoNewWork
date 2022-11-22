package main

import (
	"fmt"
	"grpc-test-chat/chat21/b"
)

func main() {

	ctx := b.New(1000, "error code")
	vt := ctx.Valuer.Get()
	fmt.Printf("%+v \n", vt)
}
