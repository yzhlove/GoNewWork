package main

import "go.uber.org/zap"

func main() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	log.Info("test zap log", zap.String("msg", "main function"))

}
