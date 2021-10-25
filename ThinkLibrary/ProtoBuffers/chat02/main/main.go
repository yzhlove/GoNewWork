package main

import (
	"fmt"
	"think-library/ProtoBuffers/chat02"
)

func main() {

	user := &chat02.User{Id: 12138, Tx: 111, Ax: 222}
	data, err := user.Marshal()
	if err != nil {
		panic(err)
	}

	if len(data) != user.Size() {
		fmt.Println("failed.")
	}

	duser := &chat02.User{}
	if err := duser.Unmarshal(data); err != nil {
		panic(err)
	}

	fmt.Println("duser => ", duser)

}
