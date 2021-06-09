package main

//go:generate msgp -tests=false -io=false

type Stu struct {
	Name     string
	Age      uint32
	Birthday string
}
