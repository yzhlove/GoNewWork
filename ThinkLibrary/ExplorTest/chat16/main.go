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

	logger.Info("zap log style => ", rec.Marshal()...)
}

type InterSource interface {
	InterCommon
	Equal(src InterSource) bool
}

type InterCommon interface {
	Id() int
	Marshal() []zap.Field
}

type Energy struct {
	ID     uint32
	Ext    string
	Handle string
}

func (e Energy) Id() int {
	return 1
}

func (e Energy) Marshal() []zap.Field {
	fields := make([]zap.Field, 0, 3)
	fields = append(fields, zap.Int("ID", e.Id()))
	fields = append(fields, zap.String("Ext", e.Ext))
	fields = append(fields, zap.String("Handle", e.Handle))
	return fields
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

func (r Res) Marshal() []zap.Field {
	fields := make([]zap.Field, 0, 2)
	fields = append(fields, zap.Int("ID", r.Id()))
	fields = append(fields, zap.String("Act", r.Act))
	return fields
}

type Record struct {
	src     InterSource
	changes []InterCommon
}

func (r Record) Marshal() []zap.Field {
	fields := make([]zap.Field, 0, 32)
	fields = append(fields, r.src.Marshal()...)
	for _, v := range r.changes {
		fields = append(fields, v.Marshal()...)
	}
	return fields
}
