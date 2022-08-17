package main

import (
	"go.uber.org/zap"
)

func main() {

	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	rec := Record{
		ID:  1001,
		Str: "what are you doing!",
		Ext: "mead in chain 1949-10-01",
	}

	log.Info("this is log style", zap.Any("rec", rec))

}

type Record struct {
	ID  uint32
	Str string
	Ext string
}
