package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
	"net/url"
	"time"
	"trace-chat/tracing"
)

func main() {

	stat := make(chan struct{})

	go startService(stat)

	trace, c, err := tracing.NewOpenTrace("trace-chat03-test-example")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	opentracing.SetGlobalTracer(trace)

	span := trace.StartSpan("begin-chat03")
	span.SetTag("start", "-----> next")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	s := formatString(ctx, "gogogo")
	s = formatString(ctx, s)
	printString(ctx, s)

	time.Sleep(time.Second)
	close(stat)
	fmt.Println("over.")
}

func formatString(ctx context.Context, str string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "<format-string>")
	defer span.Finish()

	v := url.Values{}
	v.Set("event", str)
	url := "http://localhost:8081/format?" + v.Encode()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := tracing.HttpDo(req)
	if err != nil {
		panic(err)
	}

	s := string(resp)
	span.LogFields(
		log.String("event-func", "http-format"),
		log.String("event-value", s))
	return s
}

func printString(ctx context.Context, str string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "<format-print>")
	defer span.Finish()

	v := url.Values{}
	v.Set("event", str)
	url := "http://localhost:8082/publish?" + v.Encode()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	if _, err = tracing.HttpDo(req); err != nil {
		panic(err)
	}
}

func startService(status chan struct{}) {

	go func() {
		formatService()
	}()

	go func() {
		printService()
	}()
	<-status
}

func formatService() {
	http.HandleFunc("/format", func(writer http.ResponseWriter, request *http.Request) {
		s := request.FormValue("event")
		writer.Write([]byte(fmt.Sprintf("hello:%s", s)))
	})
	fmt.Println(http.ListenAndServe(":8081", nil))
}

func printService() {
	http.HandleFunc("/publish", func(writer http.ResponseWriter, request *http.Request) {
		s := request.FormValue("event")
		println(s)
	})
	fmt.Println(http.ListenAndServe(":8082", nil))
}
