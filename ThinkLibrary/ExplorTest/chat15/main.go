package main

import (
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
	rec = rec
	encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)
	logger.Info("zap log style", rec.MarshalLog()...)

}

type Record struct {
	ID     uint32
	Ext    string
	Handle string
}

func (r Record) MarshalLog() []zap.Field {
	return []zap.Field{
		zap.Uint32("ID", r.ID),
		zap.String("Ext", r.Ext),
		zap.String("Handle", r.Handle),
	}
}
