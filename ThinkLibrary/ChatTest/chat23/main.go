package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var a = make(map[string]uint32)
	a["A"] = 100
	a["B"] = 1000
	a["C"] = 10000

	data, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	var m Map
	if err := m.Decode(string(data)); err != nil {
		panic(err)
	}

	fmt.Printf("%+v \n", m)

}

type Map map[string]uint32

func (m *Map) Decode(str string) error {
	if len(str) > 0 {
		var t = make(map[string]uint32)
		if err := json.Unmarshal([]byte(str), &t); err != nil {
			return err
		}
		*m = t
	}
	return nil
}
