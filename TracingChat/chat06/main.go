package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"net/http"
	"time"
	"trace-chat/tracing"
)

func main() {

	stat := make(chan struct{})

	go serverListener(stat)
	time.Sleep(time.Second)

	trace, clos, _ := tracing.NewOpenTrace("chat06-test-trace-example")
	defer clos.Close()

	span := trace.StartSpan("first-span")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	_ = okHttp(ctx, "http://localhost:1234/stu?event=student")
	_ = okHttp(ctx, "http://localhost:1234/teach?event=teacher")

	close(stat)

}

func openTrace(service string) opentracing.Tracer {
	cfg := config.Configuration{
		ServiceName: service,
		Sampler:     &config.SamplerConfig{Type: "const", Param: 1},
		Reporter:    &config.ReporterConfig{LogSpans: true},
	}
	trace, clo, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Errorf("new trace error:%v", err))
	}
	_ = clo
	opentracing.SetGlobalTracer(trace)
	return trace
}

func okHttp(ctx context.Context, url string) error {

	span, _ := opentracing.StartSpanFromContext(ctx, "client:"+url)
	defer span.Finish()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, req.Method)

	if err := span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header)); err != nil {
		return err
	}

	if _, err := http.DefaultClient.Do(req); err != nil {
		return err
	}

	return nil
}

var traceErr = func(e error) string {
	return fmt.Sprintf("new trace error:%v", e)
}

func stuHttpServer(w http.ResponseWriter, r *http.Request) {
	trace := openTrace("student-service")
	ctx, err := trace.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	if err != nil {
		http.Error(w, traceErr(err), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	span := trace.StartSpan("stu-service:"+r.RequestURI, ext.RPCServerOption(ctx))
	defer span.Finish()
	span.SetTag("event-service", "student-service")
	span.LogFields(
		log.String("service", "stu-http-server"),
		log.String("url", r.URL.String()))

	newContext := opentracing.ContextWithSpan(context.Background(), span)
	stuService(newContext, r.FormValue("event"))
}

func stuService(ctx context.Context, value string) {
	span, newCtx := opentracing.StartSpanFromContext(ctx, "do-stu-service")
	defer span.Finish()

	span.SetTag("event-service", "stu-func-service")
	span.SetTag("value", "do-stu:"+value)

	span.LogFields(log.String("event", value))

	stu2Service(newCtx)
}

func stu2Service(ctx context.Context) {

	span, _ := opentracing.StartSpanFromContext(ctx, "do-new-stu-service")
	defer span.Finish()

	span.SetTag("event-service", "stu2-func-service")
}

func teachHttpServer(w http.ResponseWriter, r *http.Request) {
	trace := openTrace("teacher-service")
	ctx, err := trace.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	if err != nil {
		http.Error(w, traceErr(err), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	span := trace.StartSpan("teach-service:"+r.RequestURI, ext.RPCServerOption(ctx))
	defer span.Finish()

	span.SetTag("event-service", "teacher-service")
	span.LogFields(
		log.String("service", "teach-http-server"),
		log.String("url", r.URL.String()),
		log.String("event-value", r.FormValue("event")))

}

func serverListener(stat chan struct{}) {
	go func() {
		http.HandleFunc("/stu", stuHttpServer)
		http.HandleFunc("/teach", teachHttpServer)
		log.Error(http.ListenAndServe(":1234", nil))
	}()
	<-stat
}
