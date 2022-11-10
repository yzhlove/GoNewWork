package main

import "fmt"

func main() {

	stMap := make(map[St]int)
	stMap[St{
		A: 0,
		B: 0,
	}] = 1

	stMap[St{
		A: 1,
		B: 2,
	}] = 2

	stMap[St{
		A: 2,
		B: 1,
	}] = 3

	stMap[St{
		A: 0,
		B: 0,
	}] = 4

	a := St{
		A: 0,
		B: 0,
	}

	b := St{
		A: 1,
		B: 2,
	}

	c := St{
		A: 2,
		B: 1,
	}

	fmt.Println(stMap[a])
	fmt.Println(stMap[b])
	fmt.Println(stMap[c])

}

type St struct {
	A, B int
}
