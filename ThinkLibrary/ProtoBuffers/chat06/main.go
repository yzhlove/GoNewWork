package main

import "fmt"

func main() {

	Init("127.0.0.1:6379", 1)

	stu := &Stu{Name: "abcdefgh", Age: 20000001, Birthday: "protoc-gen-gogofaster"}

	user := &User{Name: "abcdefgh", Age: 20000001, Birthday: "protoc-gen-gogofaster"}

	stuData, _ := stu.MarshalMsg(nil)
	fmt.Println("stu length => ", len(stuData))

	userData, _ := user.Marshal(nil)
	fmt.Println("user length => ", len(userData))

	conn := Get().Get()

	conn.Do("SET", "Test_Stu_Save", stuData)
	conn.Do("SET", "Test_User_Save", userData)

}
