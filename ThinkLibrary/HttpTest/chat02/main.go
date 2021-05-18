package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	s := &server{}
	a := &A{server: s}
	b := &B{server: s}

	http.Handle("/A", a)
	http.Handle("/B", b)
	log.Fatal(http.ListenAndServe(":1234", nil))

}

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("first server")
}

type A struct {
	*server
}

func (a *A) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.server.ServeHTTP(w, r)
	fmt.Println("a server")
}

type B struct {
	*server
}

func (b *B) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b.server.ServeHTTP(w, r)
	fmt.Println("b server")
}
