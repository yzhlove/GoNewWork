package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
	"net/url"
	"sync/atomic"
	"trace-chat/tracing"
)

func main() {
	stat := make(chan struct{})
	go startService(stat)

	trace, clo, err := tracing.NewOpenTrace("chat04-test-example")
	if err != nil {
		panic(err)
	}
	defer clo.Close()

	opentracing.SetGlobalTracer(trace)
	span := trace.StartSpan("first-span")
	span.SetTag("chat04", "hello tracing")
	span.SetBaggageItem("event", "trace-value")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	str := formatClient(ctx, "hello string")
	s := formatClient(ctx, str)
	printClient(ctx, s)
}

func startService(stat chan struct{}) {
	go formatService()
	go printService()
	<-stat
}

func formatClient(ctx context.Context, str string) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "string-format-client")
	defer span.Finish()

	vals := url.Values{}
	vals.Set("event", str)
	url := "http://localhost:8081/format?" + vals.Encode()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, http.MethodGet)
	span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	resp, err := tracing.HttpDo(req)
	if err != nil {
		ext.LogError(span, err)
		panic(err)
	}

	span.LogFields(log.String("event", "string-format"), log.String("value", string(resp)))

	return string(resp)
}

func printClient(ctx context.Context, str string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "string-publish-client")
	defer span.Finish()

	vals := url.Values{}
	vals.Set("event", str)
	url := "http://localhost:8082/publish?" + vals.Encode()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, http.MethodGet)
	if err := span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header)); err != nil {
		panic(err)
	}

	_, err = tracing.HttpDo(req)
	if err != nil {
		ext.LogError(span, err)
		panic(err)
	}
}

func formatService() {
	trace, c, err := tracing.NewOpenTrace("string-format-service")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var count uint32

	http.HandleFunc("/format", func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			atomic.AddUint32(&count, 1)
		}()

		ctx, err := trace.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(request.Header))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		span := trace.StartSpan("<format-service>", ext.RPCServerOption(ctx))
		defer span.Finish()
		//form baggage data
		event := span.BaggageItem("event")
		if len(event) == 0 {
			event = "no data"
		}
		evt := request.FormValue("event")
		str := fmt.Sprintf("event:%s - request:%s", event, evt)
		span.LogFields(log.String("service", "string-format"), log.String("value", str))
		writer.Write([]byte(str))
	})
	log.Error(http.ListenAndServe(":8081", nil))
}

func printService() {
	trace, clo, err := tracing.NewOpenTrace("print-format-service")
	if err != nil {
		panic(err)
	}
	defer clo.Close()

	var count uint32

	http.HandleFunc("/publish", func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			atomic.AddUint32(&count, 1)
		}()

		ctx, err := trace.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(request.Header))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		span := trace.StartSpan("<publish-service>", ext.RPCServerOption(ctx))
		defer span.Finish()

		event := span.BaggageItem("event")
		evt := request.FormValue("event")
		println(event, "-", evt)
	})
	log.Error(http.ListenAndServe(":8082", nil))
}
