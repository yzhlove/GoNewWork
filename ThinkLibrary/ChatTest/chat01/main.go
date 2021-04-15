package main

import "fmt"

func main() {

	fmt.Println("x -> ", regret())

}

func regret() (x int) {

	x = 5

	defer func() {
		fmt.Println("x = ", x)
	}()

	x = 10
	return
}
