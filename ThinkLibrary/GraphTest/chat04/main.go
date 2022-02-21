package main

import (
	"fmt"
	"sort"
)

func main() {

	a := ss{1, 5, 30, 55, 66, 100}
	sort.Sort(a)
	fmt.Println("a => ", a)
	search := 0

	idx := sort.Search(len(a), func(i int) bool {
		return a[i] <= search
	})

	fmt.Println("idx => ", idx)
}

type ss []int

func (s ss) Len() int {
	return len(s)
}

func (s ss) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ss) Less(i, j int) bool {
	return s[i] > s[j]
}
