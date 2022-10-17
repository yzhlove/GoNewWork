package main

import "fmt"

func main() {

	x, ok := Sum("hello ", "world")
	fmt.Println(x, ok)

}

func Add[T int | string](a, b T) (T, bool) {

	return a + b, true
}

func Sum(a, b string) (string, bool) {

	return Add[string](a, b)
}