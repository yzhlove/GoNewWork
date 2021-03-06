package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"grpc-test-chat/chat06/contions/anytest/echo"
)

func main() {
	encodeAnddecode()
}

func encodeAnddecode() {
	user := &echo.User{Name: "yzh", Age: 132, Password: []byte("hello world")}

	fmt.Println("toUserFormat -> \n", proto.MarshalTextString(user))
	fmt.Println("User -> ", user.String())

	userData, _ := ptypes.MarshalAny(user)

	resp := &echo.Resp{Inter: userData}
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

	if ptypes.Is(newResp.Inter, &echo.User{}) {
		nnewUser := &echo.User{}
		ptypes.UnmarshalAny(newResp.Inter, nnewUser)
		fmt.Println("nn user =====> ", nnewUser)
	}

	if err := proto.Unmarshal(newResp.Inter.Value, newUser); err != nil {
		panic(err)
	}

	fmt.Println("newUser -> ", newUser.String())

}
