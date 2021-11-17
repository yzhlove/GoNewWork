package main

import "fmt"

func main() {

	a := make(map[string]int, 4)
	a["a"] = 1
	a["b"] = 1
	a["c"] = 1
	a["d"] = 1
	a["aa"] = 2
	a["bb"] = 3
	a["cc"] = 4
	a["dd"] = 4

	for k, v := range a {
		if v == 1 {
			delete(a, k)
		}
	}

	fmt.Println("map => ", a)

}
