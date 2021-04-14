package main

import "trace-chat/tracing"

func main() {
	trace, closer, err := tracing.NewTrace("chat02-test-example")
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	trace = trace

}
