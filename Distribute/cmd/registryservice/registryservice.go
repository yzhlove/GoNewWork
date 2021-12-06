package main

import (
	"context"
	"distribute/registry"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	http.Handle("/services", registry.Service{})

	ctx, cancel := context.WithCancel(context.Background())

	var server http.Server
	server.Addr = registry.ServerPort

	waiting := make(chan os.Signal, 1)
	signal.Notify(waiting)

	go func() {
		log.Println(server.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("等待接收退出信号.")
		<-waiting
		if err := server.Shutdown(ctx); err != nil {
			fmt.Println("注册服务关闭失败,reason:", err)
		}
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("注册服务已经停止。")
}
