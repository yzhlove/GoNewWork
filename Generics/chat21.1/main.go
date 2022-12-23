package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request:", r.RemoteAddr)
		time.Sleep(time.Second * 5)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("yoxi!"))
	})
	if err := http.ListenAndServe(":7887", nil); err != nil {
		panic(err)
	}
}
