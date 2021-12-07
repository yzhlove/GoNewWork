package service

import (
	"context"
	"distribute/registry"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func Start(ctx context.Context, host, port string, reg registry.Registration, handle func()) (context.Context, error) {
	handle()
	ctx = startService(ctx, reg.ServiceName, host, port)
	if err := registry.RegistryService(reg); err != nil {
		return nil, err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {

	ctx, cancel := context.WithCancel(ctx)

	var server http.Server
	server.Addr = host + ":" + port

	waiting := make(chan os.Signal, 1)
	signal.Notify(waiting)
	go func() {
		fmt.Println(server.ListenAndServe())
		if err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port)); err != nil {
			log.Println("deregister error:", err)
		}
		cancel()
	}()

	go func() {
		fmt.Println("等待接收退出信号。")
		<-waiting
		if err := server.Shutdown(ctx); err != nil {
			fmt.Println(serviceName+"stop service:reason ", err)
		}
		if err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port)); err != nil {
			log.Println("deregister error:", err)
		}
		cancel()
	}()

	return ctx
}
