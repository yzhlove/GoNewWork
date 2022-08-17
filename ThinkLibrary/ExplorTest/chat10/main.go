package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"reflect"
)

func main() {

	rec := Record{
		Source: EnergySystem{
			ID:         1001,
			Ext:        "source energy buy handle",
			HandleName: "energy_buy_req",
			Value:      100,
		},
		Change: []InterChange{
			Res{
				ID:     1,
				Act:    "res_apply",
				Cur:    1,
				Change: 10,
			},
			Res{
				ID:     2,
				Act:    "res_exit",
				Cur:    10,
				Change: 20,
			},
			Res{
				ID:     3,
				Act:    "res_settle",
				Cur:    20,
				Change: 30,
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

type (
	InterSource interface {
		Equal(source InterSource) bool
		InterChange
	}
	InterChange interface {
		Id() int
		fmt.Stringer
	}
)

type EnergySystem struct {
	ID         uint32
	Ext        string
	HandleName string
	Value      int32
}

func (e EnergySystem) Id() int {
	return 1
}

func (e EnergySystem) Equal(src InterSource) bool {
	if x, ok := src.(EnergySystem); ok {
		return reflect.DeepEqual(e, x)
	}
	return false
}

func (e EnergySystem) String() string {
	return "source.energySystem"
}

type Res struct {
	ID     uint32
	Act    string
	Cur    int32
	Change int32
}

func (r Res) Id() int {
	return 1
}

func (r Res) String() string {
	return "res.change"
}

type Record struct {
	Source InterSource
	Change []InterChange
}

func (r Record) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddReflected("src", r.Source)
	enc.AddReflected("changes", r.Change)
	return nil
}
