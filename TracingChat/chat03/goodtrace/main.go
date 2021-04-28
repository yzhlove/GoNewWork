package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"
	"trace-chat/tracing"
)

func main() {

	stat := make(chan struct{})
	go startService(stat)
	start()
	time.Sleep(time.Second)
	close(stat)
	fmt.Println("ok ...")
}

func start() {
	t, c, err := tracing.NewOpenTrace("chat03-new-trace-example")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	opentracing.SetGlobalTracer(t)
	span := t.StartSpan("begin-chat03-span")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	var s string
	s, ctx = formatString(ctx, "gogooogo")
	s, ctx = formatString(ctx, s)
	printString(ctx, s)
	s, ctx = formatString(ctx, s)
	printString(ctx, s)
}

func formatString(ctx context.Context, str string) (string, context.Context) {
	span, newContext := opentracing.StartSpanFromContext(ctx, "{format-string}")
	defer span.Finish()

	v := url.Values{"event": []string{str}}
	u := "http://localhost:8081/format?" + v.Encode()
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		panic(err)
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, u)
	ext.HTTPMethod.Set(span, http.MethodGet)

	if err := span.Tracer().Inject(span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header)); err != nil {
		panic(err)
	}

	data, err := tracing.HttpDo(req)
	if err != nil {
		ext.LogError(span, err)
		panic(err)
	}

	span.LogFields(log.String("event", "{string-format}"),
		log.String("value", string(data)))
	return string(data), newContext
}

func printString(ctx context.Context, str string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "{print-string}")
	defer span.Finish()

	v := url.Values{"event": []string{str}}
	u := "http://localhost:8082/publish?" + v.Encode()

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		panic(err)
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, u)
	ext.HTTPMethod.Set(span, http.MethodGet)

	if err := span.Tracer().Inject(span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header)); err != nil {
		panic(err)
	}

	if _, err := tracing.HttpDo(req); err != nil {
		ext.LogError(span, err)
		panic(err)
	}

}

func formatService() {
	trace, c, err := tracing.NewOpenTrace("<format-service>")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var count int32

	http.HandleFunc("/format", func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			atomic.AddInt32(&count, 1)
		}()

		ctx, err := trace.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(request.Header))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		span := trace.StartSpan("<format-handle>", ext.RPCServerOption(ctx))
		defer span.Finish()
		if r := atomic.LoadInt32(&count); r >= 2 {
			http.Error(writer, "request server", http.StatusBadRequest)
			ext.LogError(span, errors.New("request must to many server"))
			return
		}

		str := request.FormValue("event")
		str = fmt.Sprintf("hello,%s service", str)
		span.LogFields(
			log.String("service-name", "format-service"),
			log.String("value", str))
		writer.Write([]byte(str))
	})
	log.Error(http.ListenAndServe(":8081", nil))
}

func printService() {
	trace, c, err := tracing.NewOpenTrace("<publish-service>")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	http.HandleFunc("/publish", func(writer http.ResponseWriter, request *http.Request) {
		ctx, err := trace.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(request.Header))
		if err != nil {
			panic(err)
		}
		span := trace.StartSpan("<publish-handle>", ext.RPCServerOption(ctx))
		defer span.Finish()

		println(request.FormValue("event"))
	})
	log.Error(http.ListenAndServe(":8082", nil))
}

func startService(stat chan struct{}) {
	go func() {
		formatService()
	}()
	go func() {
		printService()
	}()
	<-stat
}
