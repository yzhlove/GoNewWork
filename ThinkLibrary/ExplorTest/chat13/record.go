package main

//go:generate stringer -type=status -output=record_string.go

type status int

const (
	none status = iota
	res
	energy
)
