package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var dirPath = ""

// 创建一个静态服务器
func main() {
	fs := http.FileServer(http.Dir("."))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("=== :{%s}", r.URL.Path)
		if _, err := os.Stat(r.URL.Path); err != nil {
			log.Printf("--- :{%s}", r.URL.Path)
			http.ServeFile(w, r, "/index.html")
		} else {
			fs.ServeHTTP(w, r)
		}
	})
	fmt.Println("listing on 8001.")
	if err := http.ListenAndServe(":8001", nil); err != nil {
		panic(err)
	}
}
