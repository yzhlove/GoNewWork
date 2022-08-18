package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {

	rec := Record{
		Base: Base{
			ID:  1001,
			Ext: "base.ext",
		},
		Act:    "record.act",
		Handle: "record.handle",
		Value:  1314521,
	}

	encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)
	logger.Info("zap log style", zap.Inline(rec))
	fmt.Println()
	logger.Info("zap log style", zap.Any("record", rec))
}

type Base struct {
	ID  int
	Ext string
}

func (b Base) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("ID", b.ID)
	enc.AddString("Ext", b.Ext)
	return nil
}

type Record struct {
	Base
	Act    string
	Handle string
	Value  uint32
}

func (r Record) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if err := r.Base.MarshalLogObject(enc); err != nil {
		return err
	}
	enc.AddString("Act", r.Act)
	enc.AddString("Handle", r.Handle)
	enc.AddUint32("Value", r.Value)
	return nil
}
