package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//a := map[int]int{3: 10, 4: 20, 5: 40}
	//data, _ := json.Marshal(a)
	//fmt.Println("json text:", string(data))
	//
	//var dec json.RawMessage
	//if err := json.Unmarshal(data, &dec); err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("dec => ", string(dec), reflect.TypeOf(dec))

	test()

}

func test() {

	ss := []Dept{{100, 200}, {1000, 2000}, {330, 440}, {222, 333}, {444, 555}}

	data, _ := json.Marshal(ss)
	fmt.Println("marshal => ", string(data))

	var a Depts
	if err := a.Decode(string(data)); err != nil {
		panic(err)
	}

	fmt.Println("a --> ", a)

}

type Dept struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Depts []Dept

func (d *Depts) Decode(str string) error {
	if len(str) > 0 {
		return json.Unmarshal([]byte(str), d)
	}
	return nil
}
