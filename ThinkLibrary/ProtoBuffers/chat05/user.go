package main

//go:generate msgp -tests=false -io=false

type User struct {
	Name     string `msg:"n,omitempty"`
	Age      uint32 `msg:"a,omitempty"`
	Birthday string `msg:"b,omitempty"`
}
