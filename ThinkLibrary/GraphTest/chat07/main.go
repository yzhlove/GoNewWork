package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var a map[int][]int
	a = make(map[int][]int, 4)

	for i := 1; i <= 4; i++ {
		x := rand.Intn(10) + 1
		for k := 0; k <= x; k++ {
			a[i] = append(a[i], rand.Intn(100))
		}
	}

	show := func() {
		for k := range a {
			fmt.Println("==> ", a[k])
		}
	}

	show()

	//for k := range a {
	//	sort.Slice(a[k], func(i, j int) bool {
	//		return a[k][i] < a[k][j]
	//	})
	//}

	for k, v := range a {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		a[k] = v
	}

	fmt.Println("---------------------------")

	show()

}
