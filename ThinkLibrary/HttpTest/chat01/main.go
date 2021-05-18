package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	s := &server{}
	http.HandleFunc("/A", s.A)
	http.HandleFunc("/B", s.B)
	log.Fatal(http.ListenAndServe(":1234", nil))
}

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fist server http")
}

func (s *server) A(w http.ResponseWriter, r *http.Request) {
	s.ServeHTTP(w, r)
	fmt.Println("a server http")
}

func (s *server) B(w http.ResponseWriter, r *http.Request) {
	s.ServeHTTP(w, r)
	fmt.Println("b server http")
}
