package main

import "fmt"

type sourcer interface {
	Id() int
	Equal(src sourcer) bool
	fmt.Stringer
}

type srcbase struct {
	ID  int    `json:"id"`
	Ext string `json:"ext,omitempty"`
}

func (s srcbase) Id() int {
	return int(srcNone)
}

func (s srcbase) String() string {
	return fmt.Sprintf("ID:%d Ext:%s", s.ID, s.Ext)
}

func (s srcbase) equal(src srcbase) bool {
	return s.ID == src.ID && s.Ext == src.Ext
}

type srcTyp int

const (
	srcNone srcTyp = iota
	srcItemDrop
)
