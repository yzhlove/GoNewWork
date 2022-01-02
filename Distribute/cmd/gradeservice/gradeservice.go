package main

import (
	"context"
	"distribute/grades"
	"distribute/registry"
	"distribute/service"
	"fmt"
	stlog "log"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("%s:%s", host, port)

	ctx, err := service.Start(context.Background(), host, port, registry.Registration{
		ServiceName: registry.GradeServcie,
		ServiceURL:  fmt.Sprintf("http://%s", serviceAddress),
	}, grades.RegistryHandlers)

	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("shutting down grade service ")
}
