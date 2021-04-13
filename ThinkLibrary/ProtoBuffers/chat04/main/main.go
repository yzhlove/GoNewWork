package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"think-library/ProtoBuffers/chat04"
)

func main() {

	user := &chat04.User{}

	for i := 100; i < 10000; i++ {
		user.Ids = append(user.Ids, int32(i))
	}

	data, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println("size => ", len(data))
	//21612

}
