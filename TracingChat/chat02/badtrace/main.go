package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"trace-chat/tracing"
)

func main() {
	trace, closer, err := tracing.NewOpenTrace("chat02-test-example")
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	span := trace.StartSpan("begin-trace-span")
	//test1(span)

	//test2(span)

	test3(span)
}

func test1(span opentracing.Span) {

	formatString(span, "abc")
	formatString(span, "123")
	formatString(span, "cccc")
	formatString(span, "ffff")

	printHello(span, "hello world")
	printHello(span, "what are you doing")
	span.Finish()
}

func test2(span opentracing.Span) {

	formatString_v2(span, "aaa")
	formatString_v2(span, "bbb")
	formatString_v2(span, "ccc")
	formatString_v2(span, "ddd")
	formatString_v2(span, "123")
	formatString_v2(span, "fff")

	print_v2(span, "hello-world")
	print_v2(span, "what are you doing")

	span.Finish()
	/*
		2021/04/27 15:43:05 debug logging disabled
		2021/04/27 15:43:05 Initializing logging reporter
		2021/04/27 15:43:05 debug logging disabled
		2021/04/27 15:43:05 Reporting span 1c388923e41adeb8:1c388923e41adeb8:0000000000000000:1
		2021/04/27 15:43:05 Reporting span 7963ed13cdb88835:7963ed13cdb88835:0000000000000000:1
		2021/04/27 15:43:05 Reporting span 14522367892dfc69:14522367892dfc69:0000000000000000:1
		2021/04/27 15:43:05 Reporting span 02876b8a3d9143d4:02876b8a3d9143d4:0000000000000000:1
		2021/04/27 15:43:05 Reporting span 7292237586379c1f:7292237586379c1f:0000000000000000:1
		2021/04/27 15:43:05 Reporting span 753a8350573f92ac:753a8350573f92ac:0000000000000000:1
		hello-world
		2021/04/27 15:43:05 Reporting span 4ac70067e75bcbf9:4ac70067e75bcbf9:0000000000000000:1
		what are you doing
		2021/04/27 15:43:05 Reporting span 78f0966f836b8fbb:78f0966f836b8fbb:0000000000000000:1
		2021/04/27 15:43:05 Reporting span 1ff59ec578e85de9:1ff59ec578e85de9:0000000000000000:1
	*/
}

func test3(span opentracing.Span) {
	formatString_v3(span, "aaa")
	formatString_v3(span, "bbb")

	print_v3(span, "hello-world")
	print_v3(span, "what are you doing")
	span.Finish()
}

func formatString(span opentracing.Span, str string) string {
	s := fmt.Sprintf("Hello,%s", str)
	span.LogFields(
		log.String("func-name", "string-format"),
		log.String("func-value", s))
	return s
}

func printHello(span opentracing.Span, str string) {
	println(str)
	span.LogKV("log-kv-func", "print-hello")
}

func formatString_v2(root opentracing.Span, str string) string {
	span := root.Tracer().StartSpan("<format-string>")
	defer span.Finish()

	s := fmt.Sprintf("Hello,%s", str)
	span.LogFields(
		log.String("func-name", "string-format"),
		log.String("func-value", s))
	return s
}

func print_v2(root opentracing.Span, str string) {
	span := root.Tracer().StartSpan("<format-print>")
	defer span.Finish()

	println(str)
	span.LogKV("log-kv-func", str)
}

func formatString_v3(root opentracing.Span, str string) string {
	span := root.Tracer().StartSpan("<string-format>", opentracing.ChildOf(root.Context()))
	defer span.Finish()
	s := fmt.Sprintf("Hello,%s", str)
	span.LogFields(
		log.String("func-name", "string-format"),
		log.String("func-value", s))
	return s
}

func print_v3(root opentracing.Span, str string) {
	span := root.Tracer().StartSpan("<string-print>", opentracing.ChildOf(root.Context()))
	defer span.Finish()

	println(str)
	span.LogKV("log-kv-func", str)
}
