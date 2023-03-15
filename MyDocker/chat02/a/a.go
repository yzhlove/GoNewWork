package a

import "my-docker/chat02/b"

type A struct {
	ctx *b.B
}

func (a *A) Init(bb *b.B) {
	a.ctx = bb
}

func (a *A) Run() string {
	return "start a is running..."
}
