package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const value = `{"0#src.id": 1001, "0#src.ext": "energy_ext", "0#src.handle": "energy_buy_req", "0#changes.0#change.id": 100, "0#changes.0#change.act": "changes_1", "0#changes.1#change.id": 200, "0#changes.1#change.act": "changes_2", "1#src.id": 2001, "1#src.ext": "energy_ext", "1#src.handle": "energy_buy_req", "1#changes.0#change.id": 2100, "1#changes.0#change.act": "changes_11", "1#changes.1#change.id": 2200, "1#changes.1#change.act": "changes_21", "2#src.id": 3001, "2#src.ext": "energy_ext", "2#src.handle": "energy_buy_req", "2#changes.0#change.id": 3100, "2#changes.0#change.act": "changes_13", "2#changes.1#change.id": 3200, "2#changes.1#change.act": "changes_23"}`

func main() {

	var data = make(map[string]interface{})

	if err := json.Unmarshal([]byte(value), &data); err != nil {
		panic(err)
	}

	for k, v := range data {
		fmt.Println(k, " = ", v, " type=>", reflect.TypeOf(v).Name())
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("bytes => ", string(bytes))

	fmt.Println("-------------------------------------")
	test()

}

type A struct {
	Id  int
	Ext string
}

type B struct {
	A
	Handle string
}

func test() {

	str1 := "0#a.1#b.c"

	str2 := "0#a@1#b.c"

	str3 := "0#src@1#B.A.Id"

	str4 := "src@1.B.A.Id@2"

	str5 := "group.1 "

	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c", str1[i])
	}

}
