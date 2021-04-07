package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"think-library/chat01"
)

func main() {

	user := &chat01.User{Id: 150, Ax: 150, Tx: 150}
	data, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", data)

}
