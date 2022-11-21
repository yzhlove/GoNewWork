package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
)

func main() {

	resp := &Resp_ErrDetail{ErrDetail: "----> error string"}

	result := &Resp{Err: resp}
	msg, err := proto.Marshal(result)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(msg))

	resp2 := &Resp_ErrCode{ErrCode: 1024}

	result.Err = resp2
	msg, err = proto.Marshal(result)
	if err != nil {
		panic(err)
	}

	fmt.Println(msg)
}
