package main

import (
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
		changes: []InterCommon{
			Res{ID: 2001, Act: "res_1"},
			Res{ID: 2002, Act: "res_2"},
			Res{ID: 2003, Act: "res_3"},
		},
	}

	encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)

	logger.Info("inline energy =>>> ", zap.Inline(rec.src))
	logger.Info("inline changes =>> ", zap.Inline(rec.changes))
	logger.Info("inline record ---> ", zap.Inline(rec))

}

type InterSource interface {
	InterCommon
	Equal(src InterSource) bool
}

type InterCommon interface {
	Id() int
	zapcore.ObjectMarshaler
}

type Energy struct {
	ID     uint32
	Ext    string
	Handle string
}

func (e Energy) Id() int {
	return 1
}

func (e Energy) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("ID", e.Id())
	enc.AddString("Ext", e.Ext)
	enc.AddString("Handle", e.Handle)
	return nil
}

func (e Energy) Equal(src InterSource) bool {
	if x, ok := src.(Energy); ok {
		return x.ID == e.ID
	}
	return false
}

type Res struct {
	ID  uint32
	Act string
}

func (r Res) Id() int {
	return 2
}

func (r Res) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddUint32("ID", r.ID)
	enc.AddString("Act", r.Act)
	return nil
}

type changes []InterCommon

func (c changes) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	for _, k := range c {
		if err := k.MarshalLogObject(enc); err != nil {
			return err
		}
	}
	return nil
}

type Record struct {
	src InterSource
	changes
}

func (r Record) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if err := r.src.MarshalLogObject(enc); err != nil {
		return err
	}
	if err := r.changes.MarshalLogObject(enc); err != nil {
		return err
	}
	return nil
}
