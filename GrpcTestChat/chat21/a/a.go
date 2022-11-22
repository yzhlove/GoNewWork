package a

import "grpc-test-chat/chat21/b"

type Value struct {
	ctx *b.Context
	no  uint32
}

func (v *Value) New(ctx *b.Context, no uint32) {
	if v == nil {
		*v = Value{ctx: ctx, no: no}
	}
}

func (v *Value) Get() *Value {
	return v
}
