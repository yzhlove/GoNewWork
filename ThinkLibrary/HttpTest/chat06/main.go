package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {

	req, _ := http.NewRequest(http.MethodPost, "localhost", strings.NewReader("hello world"))
	req2 := req.Clone(req.Context())

	contents, _ := io.ReadAll(req.Body)
	contents2, _ := io.ReadAll(req2.Body)

	log.Printf("contents => %s", string(contents))
	log.Printf("contents2 => %s", string(contents2))
}
