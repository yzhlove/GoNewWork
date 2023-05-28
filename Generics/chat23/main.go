package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "a b 'd e'"
	fmt.Println("values => ", strings.SplitN(str, " ", 3))

}
