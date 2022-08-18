package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	r := record{Id: res, Name: "what"}
	fmt.Println(r)

	data, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	fmt.Println("record string => ", string(data))

}

type record struct {
	Id   status
	Name string
}
