package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

//sha256 by file

func main() {
	fmt.Println(fileSha256("ChatTest/chat13/main.go"))
}

func fileSha256(file string) string {

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	h256 := sha256.New()
	io.Copy(h256, f)

	return fmt.Sprintf("%x", h256.Sum(nil))
}
