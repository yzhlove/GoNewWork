package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	show(file)
	file.Seek(0, 0)
	if err = file.Truncate(0); err != nil {
		fmt.Println(err)
	}

	show(file)
	file.Close()
}

func show(f *os.File) {

	res, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Name()+" size:", res.Size())
}
