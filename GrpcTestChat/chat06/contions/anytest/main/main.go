package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"grpc-test-chat/chat06/contions/anytest/echo"
)

func main() {
	encodeAnddecode()
}

func encodeAnddecode() {
	user := &echo.User{Name: "yzh", Age: 132, Password: []byte("hello world")}
	userBytes, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Println("toUserFormat -> \n", proto.MarshalTextString(user))
	fmt.Println("User -> ", user.String())

	resp := &echo.Resp{Inter: &any.Any{Value: userBytes}}
	respBytes, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}

	fmt.Println("Resp -> ", resp.String())

	newUser := &echo.User{}
	newResp := &echo.Resp{}

	if err := proto.Unmarshal(respBytes, newResp); err != nil {
		panic(err)
	}

	fmt.Println("newResp -> ", newResp.String())

	fmt.Println("type url ----> ", newResp.Inter.TypeUrl)
	
	if err := proto.Unmarshal(newResp.Inter.Value, newUser); err != nil {
		panic(err)
	}

	fmt.Println("newUser -> ", newUser.String())

}
