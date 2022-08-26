package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {

	rec := record{
		src: energy{
			id:     1001,
			ext:    "energy",
			handle: "energy_req",
		},
		changes: []inter{
			res{
				id:  100,
				act: "res",
			},
			worldClass{
				id:  200,
				act: "worldClass",
			},
		},
	}

	log := logger()
	log.Info("style", zap.Inline(rec))

}

func logger() *zap.Logger {
	encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)
	return logger
}

type inter interface {
	encode(prefix string, enc zapcore.ObjectEncoder) error
}

type energy struct {
	id     int
	ext    string
	handle string
}

func (e energy) encode(prefix string, enc zapcore.ObjectEncoder) error {
	enc.AddInt("id", e.id)
	enc.AddString("ext", e.ext)
	enc.AddString("handle", e.handle)
	return nil
}

type res struct {
	id  int
	act string
}

func (r res) encode(prefix string, enc zapcore.ObjectEncoder) error {
	enc.AddInt(prefix+".res.id", r.id)
	enc.AddString(prefix+".res.act", r.act)
	return nil
}

type worldClass struct {
	id  int
	act string
}

func (r worldClass) encode(prefix string, enc zapcore.ObjectEncoder) error {
	enc.AddInt(prefix+"world.id", r.id)
	enc.AddString(prefix+"world.act", r.act)
	return nil
}

type inters []inter

type record struct {
	src     inter
	changes inters
}

func (r record) MarshalLogObject(enc zapcore.ObjectEncoder) error {

	if err := r.src.encode("src", enc); err != nil {
		return err
	}

	for _, c := range r.changes {
		if err := c.encode("change", enc); err != nil {
			return err
		}
	}
	return nil
}
