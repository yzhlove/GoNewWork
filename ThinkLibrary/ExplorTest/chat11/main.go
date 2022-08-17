package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	rec := Record{
		Source: EnergySystem{
			ID:     1001,
			Ext:    "source energy buy handle",
			Handle: "energy_buy_req",
		},
		Change: []InterChange{
			Res{
				ID:  1,
				Act: "res_apply",
			},
			Res{
				ID:  2,
				Act: "res_exit",
			},
			Res{
				ID:  3,
				Act: "res_settle",
			},
		},
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(
			zap.NewProductionEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel)

	logger := zap.New(core)

	logger.Info("zap log style", zap.Inline(rec))
	fmt.Println()
	logger.Info("zap log style", zap.Any("carrier", rec))

}

type InterSource interface {
	InterChange
	Equal(src InterSource) bool
}

type InterChange interface {
	Id() int
	zapcore.ObjectMarshaler
}

type EnergySystem struct {
	ID     uint32
	Ext    string
	Handle string
}

func (es EnergySystem) Id() int {
	return 1
}

func (es EnergySystem) Equal(src InterSource) bool {
	return false
}

func (es EnergySystem) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddUint32("ID", es.ID)
	enc.AddString("Ext", es.Ext)
	enc.AddString("handle", es.Handle)
	return nil
}

type Res struct {
	ID  uint32
	Act string
}

func (res Res) Id() int {
	return 2
}

func (res Res) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddUint32("ID", res.ID)
	enc.AddString("Act", res.Act)
	return nil
}

type Changes []InterChange

func (cs Changes) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, c := range cs {
		enc.AppendObject(c)
	}
	return nil
}

type Record struct {
	Source InterSource
	Change Changes
}

func (r Record) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddObject("src", r.Source)
	enc.AddArray("changes", r.Change)
	return nil
}

/*
1.6607539255728798e+09  info    zap log style
{"src": {"ID": 1001, "Ext": "source energy buy handle", "handle": "energy_buy_req"}, "changes": [{"ID": 1, "Act": "res_apply"}, {"ID": 2, "Act": "res_exit"}, {"ID": 3, "Act": "res_settle"}]}

1.660753925572961e+09   info    zap log style
{"carrier": {"src": {"ID": 1001, "Ext": "source energy buy handle", "handle": "energy_buy_req"}, "changes": [{"ID": 1, "Act": "res_apply"}, {"ID": 2, "Act": "res_exit"}, {"ID": 3, "Act": "res_settle"}]}}

*/

/*
1.6607539477767098e+09  info    zap log style
{"src": {"ID":1001,"Ext":"source energy buy handle","HandleName":"energy_buy_req","Value":100}, "changes": [{"ID":1,"Act":"res_apply","Cur":1,"Change":10},{"ID":2,"Act":"res_exit","Cur":10,"Change":20},{"ID":3,"Act":"res_settle","Cur":20,"Change":30}]}

1.660753947776923e+09   info    zap log style
{"carrier": {"src": {"ID":1001,"Ext":"source energy buy handle","HandleName":"energy_buy_req","Value":100}, "changes": [{"ID":1,"Act":"res_apply","Cur":1,"Change":10},{"ID":2,"Act":"res_exit","Cur":10,"Change":20},{"ID":3,"Act":"res_settle","Cur":20,"Change":30}]}}
*/
