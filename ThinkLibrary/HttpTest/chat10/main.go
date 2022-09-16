package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/japan", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(fmt.Sprintf("hello %s", time.Now().Format(time.RFC3339))))
	})
	log.Println("listening to 2345")
	http.ListenAndServe(":2345", nil)

}
