package main

import "fmt"

func main() {

	x, ok := Add[string]("hello ", "world")
	fmt.Println(x, ok)

}

func Add[T int | string](a, b T) (T, bool) {

	return a + b, true
}
