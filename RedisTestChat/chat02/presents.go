package main

//go:generate msgp -tests=false -io=false

//msgp:tuple Value
type Value struct {
	Star  int32 `msg:s,omitempty`
	Left  int32 `msg:"a,omitempty"`
	Right int32 `msg:"b,omitempty"`
}

type Values [6]Value
