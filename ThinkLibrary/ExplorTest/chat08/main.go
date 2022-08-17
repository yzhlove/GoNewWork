package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	rec := Record{
		ID:  1001,
		Str: "what are you doing!",
		Ext: "mead in chain 1949-10-01",
	}

	log.Info("this is log style", zap.Inline(rec))

}

type Record struct {
	ID  uint32
	Str string
	Ext string
}

func (r Record) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddUint32("ID", r.ID)
	enc.AddString("Str", r.Str)
	enc.AddString("Ext", r.Ext)
	enc.AddReflected("recordReflect", r)
	return nil
}
