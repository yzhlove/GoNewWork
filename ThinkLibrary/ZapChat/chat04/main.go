package main

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

type source struct {
	Id     int
	Handle string
}

type change struct {
	Id   int
	Act  string
	Desc string
}

type adapter struct {
	Source  source
	Changed []change
}

func main() {

	src := source{Id: 1, Handle: "main_line_settle_req"}
	change1 := change{Id: 1, Act: "added", Desc: "presents added"}
	change2 := change{Id: 2, Act: "deleted", Desc: "outfit deleted"}

	apt := &adapter{src, []change{change1, change2}}
	data, err := json.Marshal(apt)
	if err != nil {
		panic(err)
	}
	fmt.Println("json: ", string(data))

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	logger.Info("json carrier marshal", zap.Any("metadata", apt))

}
