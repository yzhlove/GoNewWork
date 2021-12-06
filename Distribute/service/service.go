package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func Start(ctx context.Context, srvName, host, port string, handle func()) (context.Context, error) {
	handle()
	ctx = startService(ctx, srvName, host, port)
	return ctx, nil
}

func startService(ctx context.Context, srvName, host, port string) context.Context {

	ctx, cancle := context.WithCancel(ctx)

	var server http.Server
	server.Addr = host + ":" + port

	waiting := make(chan os.Signal, 1)
	signal.Notify(waiting)
	go func() {
		fmt.Println(server.ListenAndServe())
		cancle()
	}()

	go func() {
		fmt.Println("等待接收退出信号。")
		<-waiting
		if err := server.Shutdown(ctx); err != nil {
			fmt.Println("停止服务错误:reason ", err)
		}
		cancle()
	}()

	return ctx
}
