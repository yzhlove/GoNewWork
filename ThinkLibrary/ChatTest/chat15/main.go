package main

import (
	"fmt"
	"os"
)

func main() {

	entr, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, et := range entr {
		fmt.Println("et.name:", et.Name())
		fmt.Println("et.type:", et.Type())
		fmt.Println("et.isDir()", et.IsDir())
		fmt.Println()
	}

}
