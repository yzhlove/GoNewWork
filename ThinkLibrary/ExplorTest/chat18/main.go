package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {

	rec := Record{
		src: Energy{
			ID:     1001,
			Ext:    "energy change",
			Handle: "energy_handle_req",
		},
		changes: []commoner{
			Res{ID: 2001, Act: "res_1"},
			Res{ID: 2002, Act: "res_2"},
			Res{ID: 2003, Act: "res_3"},
		},
	}

	encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)

	logger.Info("inline record =>> ", zap.Inline(rec))

}

type commoner interface {
	Id() int
	Encode(prefix string, enc zapcore.ObjectEncoder) error
}

type sourcer interface {
	commoner
	Equal(src sourcer) bool
}

type Energy struct {
	ID     uint32
	Ext    string
	Handle string
}

type Res struct {
	ID  uint32
	Act string
}

type changes []commoner

type Record struct {
	src sourcer
	changes
}

func (e Energy) Id() int {
	return 1
}

func (e Energy) Encode(prefix string, enc zapcore.ObjectEncoder) error {
	enc.AddUint32(prefix+".ID", e.ID)
	enc.AddString(prefix+".Ext", e.Ext)
	enc.AddString(prefix+".Handle", e.Handle)
	return nil
}

func (e Energy) Equal(src sourcer) bool {
	return false
}

func (r Res) Id() int {
	return 2
}

func (r Res) Encode(prefix string, enc zapcore.ObjectEncoder) error {
	enc.AddUint32(prefix+".ID", r.ID)
	enc.AddString(prefix+".Act", r.Act)
	return nil
}

func (c changes) Encode(prefix string, enc zapcore.ObjectEncoder) error {
	for k, v := range c {
		if err := v.Encode(fmt.Sprintf("%s.%d", prefix, k), enc); err != nil {
			return err
		}
	}
	return nil
}

func (r Record) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if err := r.src.Encode("src", enc); err != nil {
		return err
	}
	return r.changes.Encode("changes", enc)
}
