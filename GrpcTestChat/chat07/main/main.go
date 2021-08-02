package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"grpc-test-chat/chat07/echo"
)

func main() {
	testGogoProtoc()
}

func testGogoProtoc() {
	user := &echo.User{Name: "abc", Age: 100, Address: "China"}
	fmt.Println("1 user protoc -> \n", proto.MarshalTextString(user))

	userData, _ := types.MarshalAny(user)
	resp := &echo.Resp{Name: user.Name, Res: userData}
	fmt.Println("2 resp protoc -> \n", proto.MarshalTextString(resp))

	respData, _ := proto.Marshal(resp)

	newResp := &echo.Resp{}
	proto.Unmarshal(respData, newResp)
	fmt.Println("3 new resp protoc -> \n", proto.MarshalTextString(newResp))

	fmt.Println("====== type url ------> ", newResp.Res.TypeUrl)

	if types.Is(newResp.Res, &echo.User{}) {
		newUser := &echo.User{}
		if err := types.UnmarshalAny(newResp.Res, newUser); err != nil {
			panic(err)
		}

		fmt.Println("4 new user protoc -> \n", proto.MarshalTextString(newUser))
	} else {
		fmt.Println("4 types assert error")
	}

	newUser2 := &echo.User{}
	proto.Unmarshal(newResp.Res.Value, newUser2)
	fmt.Println("5 new user 2 protoc -> ", proto.MarshalTextString(newUser2))
}
