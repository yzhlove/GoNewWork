package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"trace-chat/tracing"
)

func main() {

	trace, c, err := tracing.NewOpenTrace("chat02-test-new-example")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	opentracing.SetGlobalTracer(trace)

	span := trace.StartSpan("begin-chat02-test")
	span.SetTag("hello-to", "<begin>")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	ss := stringFormat(ctx, "aaaaa")
	bb := stringFormat(ctx, ss)
	stringPrint(ctx, bb)

}

func stringFormat(ctx context.Context, str string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "<string-format>")
	defer span.Finish()
	s := fmt.Sprintf("hello:%s", str)
	span.LogFields(
		log.String("event", "<func>-string-format"),
		log.String("value", s))
	return s
}

func stringPrint(ctx context.Context, str string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "<string-print>")
	defer span.Finish()

	println(str)
	span.LogKV("<event-print>", str)
}
