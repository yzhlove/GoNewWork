package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	s := []string{
		"1",
		"2",
		"3",
		"1122",
		"1112",
		"1abc",
		"1acc",
		"abc",
		"abd",
		"aab",
		"acc",
	}

	for k, v := range rand.Perm(len(s)) {
		s[k], s[v] = s[v], s[k]
	}

	sort.Strings(s)
	fmt.Println(s)

}
