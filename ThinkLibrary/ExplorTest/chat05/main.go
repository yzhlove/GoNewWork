package main

import (
	"bytes"
	"encoding/json"
	"log"
)

func main() {
	u := &User{A: "aaa", B: 100}
	data, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	log.Println("marshal:", string(data))
}

type User struct {
	A string
	B int
}

func (u User) MarshalJSON() ([]byte, error) {
	ret := bytes.NewBufferString(`{"abc":"def"}`)
	return ret.Bytes(), nil
}
