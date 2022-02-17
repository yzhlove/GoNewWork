package main

import "fmt"

type Arr [2][2]int

func (a *Arr) add(index, count int) {
	if index >= 0 && len(a) > index {
		a[index][0] += count
	}
}

func main() {

	a := [2]int{1, 2}

	b := a
	fmt.Println(b)
	b[1] = 10
	fmt.Println(a, b)

	var c Arr
	fmt.Println(c)
	c.add(0, 100)
	fmt.Println(c)

	var d Arr
	d = c
	fmt.Println(d)

	d.add(1, 100)
	fmt.Println(c, d)

}
