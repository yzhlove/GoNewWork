package main

import "fmt"

func main() {
	testb()
}

func testa() (str string) {

	defer func() {
		if x := recover(); x != nil {
			fmt.Println(fmt.Sprint(x))
		}
		str = "hello world"
	}()

	panic("what are you doing")
}

func testb() {
	fmt.Println(testa())
}
