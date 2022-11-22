package b

import (
	"grpc-test-chat/chat21/c"
)

type Context struct {
	Str string
	c.Valuer
}

func New(no uint32, str string) *Context {
	ctx := &Context{Str: str}
	ctx.Valuer.New(ctx, no)
	return ctx
}
