package main

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	logger.Info("test zap logger", zap.String("msg", "a"), zap.String("msg", "b"))

	str := `{"level":"info","ts":1659853564.1813881,"caller":"chat03/main.go:13","msg":"test zap logger","msg":"a","msg":"b"}`
	decodeLog(str)
}

func decodeLog(str string) {

	data := make(map[string]interface{})

	if err := json.NewDecoder(strings.NewReader(str)).Decode(&data); err != nil {
		panic(err)
	}

	fmt.Printf("%#v", data)

}
