package main

import "fmt"

func main() {

	t := &s{aa: map[int]int{1: 1, 2: 2, 3: 3}}
	fmt.Println(t.aa)

	rt := t.export()
	fmt.Println(rt)

	rt[1] = 10
	rt[2] = 20
	rt[3] = 30
	rt[4] = 40
	rt[5] = 40
	rt[6] = 40
	rt[7] = 40
	rt[8] = 40
	rt[9] = 40
	fmt.Println(t.aa, rt)

}

type s struct {
	aa map[int]int
}

func (st *s) export() map[int]int {
	return st.aa
}
