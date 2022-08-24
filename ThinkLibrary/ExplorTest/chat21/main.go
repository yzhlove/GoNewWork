package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

func main() {
	rec := records{
		record{
			src: energy{
				id:     1001,
				ext:    "energy_ext",
				handle: "energy_buy_req",
			},
			changes: []inter{
				res{
					id:  100,
					act: "changes_1",
				},
				res{
					id:  200,
					act: "changes_2",
				},
			},
		},

		record{
			src: energy{
				id:     2001,
				ext:    "energy_ext",
				handle: "energy_buy_req",
			},
			changes: []inter{
				res{
					id:  2100,
					act: "changes_11",
				},
				res{
					id:  2200,
					act: "changes_21",
				},
			},
		},

		record{
			src: energy{
				id:     3001,
				ext:    "energy_ext",
				handle: "energy_buy_req",
			},
			changes: []inter{
				res{
					id:  3100,
					act: "changes_13",
				},
				res{
					id:  3200,
					act: "changes_23",
				},
			},
		},
	}

	log := logger()
	log.Info("zap log style >> ", zap.Inline(rec))

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

func build(strs ...string) string {
	return strings.Join(strs, ".")
}

func (e energy) encode(prefix string, enc zapcore.ObjectEncoder) error {
	enc.AddInt(build(prefix, "id"), e.id)
	enc.AddString(build(prefix, "ext"), e.ext)
	enc.AddString(build(prefix, "handle"), e.handle)
	return nil
}

type res struct {
	id  int
	act string
}

func (r res) encode(prefix string, enc zapcore.ObjectEncoder) error {
	enc.AddInt(build(prefix, "id"), r.id)
	enc.AddString(build(prefix, "act"), r.act)
	return nil
}

type inters []inter

func (is inters) encode(prefix string, enc zapcore.ObjectEncoder) error {
	for k, i := range is {
		if err := i.encode(build(prefix, fmt.Sprintf("%d#change", k)), enc); err != nil {
			return err
		}
	}
	return nil
}

type record struct {
	src     inter
	changes inters
}

type records []record

func (r records) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	for k, t := range r {
		if err := t.src.encode(fmt.Sprintf("%d#src", k), enc); err != nil {
			return err
		}
		if err := t.changes.encode(fmt.Sprintf("%d#changes", k), enc); err != nil {
			return err
		}
	}
	return nil
}
