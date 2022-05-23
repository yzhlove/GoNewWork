package main

import (
	"fmt"
	"net/http"
)

// 创建一个静态服务器
func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	fmt.Println("listing on 8000.")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
