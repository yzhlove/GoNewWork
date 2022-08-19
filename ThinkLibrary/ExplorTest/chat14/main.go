package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {

	rec := Record{
		ID:     1001,
		Ext:    "what",
		Handle: "are you doing",
	}

	encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)
	logger.Info("zap log style", zap.Inline(rec))
	fmt.Println()
	logger.Info("zap log style", zap.Any("record", rec))

}

type Record struct {
	ID     int
	Ext    string
	Handle string
}

func (r Record) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("ID", r.ID)
	enc.AddString("Ext", r.Ext)
	enc.AddString("Handle", r.Handle)
	return nil
}
