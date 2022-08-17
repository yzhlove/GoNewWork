package main

import "fmt"

type changer interface {
	Id() int
	fmt.Stringer
}

type changebase struct {
	ID  int    `json:"id"`
	Act string `json:"act,omitempty"`
}

type changeTyp int

const (
	changeNone changeTyp = iota
	changeItem
)
