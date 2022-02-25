package main

import (
	"encoding/json"
	"fmt"
)

type str struct {
	A int    `json:"a,omitempty"`
	B int    `json:"b,omitempty"`
	C string `json:"c,omitempty"`
}

func main() {

	r, err := json.Marshal(str{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("len: %d %s", len(r), string(r))

}
