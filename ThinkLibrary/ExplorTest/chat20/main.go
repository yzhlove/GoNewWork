package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

// {"src.ID": 1001, "src.Ext": "energy change", "src.Handle": "energy_handle_req",
//"changes.0.ID": 2001, "changes.0.Act": "res_1", "changes.1.ID": 2002, "changes.1.Act": "res_2", "changes.2.ID": 2003, "changes.2.Act": "res_3"}

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

type record struct {
	src     inter
	changes inters
}

func (r record) encode(prefix string, enc zapcore.ObjectEncoder) error {
	if err := r.src.encode(build(prefix, "src"), enc); err != nil {
		return err
	}
	for k, c := range r.changes {
		if err := c.encode(build(prefix, "changes", fmt.Sprintf("%d", k)), enc); err != nil {
			return err
		}
	}
	return nil
}

type records []record

func (r records) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	for k, t := range r {
		if err := t.encode(fmt.Sprintf("%d", k), enc); err != nil {
			return err
		}
	}
	return nil
}
