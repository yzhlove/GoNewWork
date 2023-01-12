package main

import (
	"generics-chat/chat20.2/log"
	"go.uber.org/zap"
)

func main() {

	log.Info("info ", zap.String("name", "abc"))

}
