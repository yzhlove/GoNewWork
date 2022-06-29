package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var data = []byte(`{"Succeed":true,"Result":"hello world!"}`)

	var resp = &struct {
		Succeed bool
		Result  string
	}{}

	if err := json.Unmarshal(data, resp); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", resp)

}
