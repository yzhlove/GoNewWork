package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {

	e := energy{
		id:     1001,
		ext:    "hello world",
		values: []string{"a", "b", "c"},
		items:  []item{{10000, 100}, {20000, 100}, {30000, 100}},
	}

	encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)

	logger.Info("inline record =>> ", zap.Inline(e))

}

type inter interface {
	encode(group string, enc zapcore.ObjectMarshaler) error
}

type item struct {
	id  uint32
	qty int32
}

type energy struct {
	id     int
	ext    string
	values []string
	items  []item
}

func (e energy) encode(group string, enc zapcore.ObjectEncoder) error {
	enc.AddInt(fmt.Sprintf("%s.%s", group, "id"), e.id)
	enc.AddString(fmt.Sprintf("%s.%s", group, "ext"), e.ext)
	for k, v := range e.values {
		enc.AddString(fmt.Sprintf("%s.%s.%d", group, "values", k), v)
	}
	for k, v := range e.items {
		enc.AddUint32(fmt.Sprintf("%s.%s.%s.%d", group, "items", "id", k), v.id)
		enc.AddInt32(fmt.Sprintf("%s.%s.%s.%d", group, "items", "qty", k), v.qty)
	}
	return nil
}

func (e energy) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return e.encode("energy", enc)
}
