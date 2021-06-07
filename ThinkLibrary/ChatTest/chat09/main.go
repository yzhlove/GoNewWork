package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//r := []int{8}
	r := []int{7, 8} //panic
	for k, v := range a {
		for _, vv := range r {
			if v == vv {
				a = append(a[:k], a[k+1:]...)
			}
		}
	}

	fmt.Println("a", a)

}
