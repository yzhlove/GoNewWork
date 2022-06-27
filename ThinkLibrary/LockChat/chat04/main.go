package main

import (
	"encoding/json"
	"log"
)

func main() {

	data := A{Name: "a_name", B: &B{}}
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	log.Println("marshal -> ", string(bytes))

}

type A struct {
	Name string
	B    *B `json:"b,omitempty"`
}

type B struct {
	N string
	S string
}
