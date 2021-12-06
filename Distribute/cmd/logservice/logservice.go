package main

import (
	"context"
	"distribute/log"
	"distribute/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distribute.log")
	host, port := "localhost", "4399"
	ctx, err := service.Start(context.Background(), "logService", host, port, log.RegisterHandles)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("服务关闭.")
}
