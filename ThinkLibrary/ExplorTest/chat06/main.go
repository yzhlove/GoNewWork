package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	ID uint32
}

type B struct {
	ID   uint32
	Desc string
}

type Shower interface {
	Show() string
}

func (a A) Show() string {
	data, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (b B) Show() string {
	data, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {

	a1 := A{ID: 1}
	a2 := A{ID: 2}
	a3 := A{ID: 2}

	b1 := B{ID: 3, Desc: "this b1"}
	b2 := B{ID: 4, Desc: "this b2"}

	data := make(map[Shower]struct{})
	data[a1] = struct{}{}
	data[a2] = struct{}{}
	data[a3] = struct{}{}
	data[b1] = struct{}{}
	data[b2] = struct{}{}

	for k := range data {
		fmt.Println(k.Show())
	}

}
