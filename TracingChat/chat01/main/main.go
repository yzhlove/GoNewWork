package main

import (
	"fmt"
	"os"
	"trace-chat/tracing"
)

func main() {
	tracer, closer, err := tracing.NewOpenTrace("frontend-test-jaeger")
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	params := os.Args[0]
	span := tracer.StartSpan("start-span-hello")
	//span.SetTag("hello", "world")

	str := fmt.Sprintf("Hello %s!", params)
	//span.LogFields(log.String("event", "string-format"), log.String("value", params))

	fmt.Println("str -> ", str)
	//span.LogKV("event", "span-logkv")
	span.Finish()
}
