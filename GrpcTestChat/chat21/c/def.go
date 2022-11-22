package c

import (
	"grpc-test-chat/chat21/a"
	"grpc-test-chat/chat21/b"
)

type Valuer interface {
	New(ctx *b.Context, no uint32)
	Get() *a.Value
}
