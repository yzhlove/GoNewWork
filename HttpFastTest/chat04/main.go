package main

import (
	"fmt"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	fmt.Println("listing on 8001.")
	if err := http.ListenAndServe(":8001", nil); err != nil {
		panic(err)
	}

}
