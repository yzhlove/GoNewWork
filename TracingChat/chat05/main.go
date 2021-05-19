package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"net/http"
	"time"
)

/*
				->E
		B->		->F
A -> 	C->		>>G
		D->		H->		J->K
				I->
*/

func main() {

	runServer()

	trace := openTrace("chat05-test-tracing")
	span := trace.StartSpan("first-main")
	span.SetTag("start-point", "chat04")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	for _, u := range []string{
		"http://localhost:1234/A?a=1000",
		"http://localhost:1234/B?b=2000",
		"http://localhost:1234/C?c=3000",
	} {
		if err := okHttp(ctx, u); err != nil {
			log.Error(fmt.Errorf("url:%s err:%v", u, err))
		}
	}

	time.Sleep(time.Second)
}

func openTrace(service string) opentracing.Tracer {
	cfg := config.Configuration{
		ServiceName: service,
		Sampler:     &config.SamplerConfig{Type: "const", Param: 1},
		Reporter:    &config.ReporterConfig{LogSpans: true},
	}
	trace, _, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Errorf("new trace error:%v", err))
	}
	opentracing.SetGlobalTracer(trace)
	return trace
}

func httpDo(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
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
	ext.HTTPMethod.Set(span, http.MethodGet)
	if err := span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header)); err != nil {
		return err
	}

	if data, err := httpDo(req); err != nil {
		return err
	} else {
		fmt.Println("response data -> ", string(data))
	}
	return nil
}

var traceErr = func(err error) string {
	return fmt.Sprintf("trace extract error:%s", err)
}

type server struct {
	trace opentracing.Tracer
}

func (s *server) serverHttp(w http.ResponseWriter, r *http.Request) {
	ctx, err := s.trace.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	if err != nil {
		http.Error(w, traceErr(err), http.StatusInternalServerError)
		return
	}

	span := s.trace.StartSpan("server:"+r.URL.String(), ext.RPCServerOption(ctx))
	defer span.Finish()

	span.LogFields(log.String("service", r.URL.String()))
	r.WithContext(opentracing.ContextWithSpan(r.Context(), span))
}

func (s *server) A(w http.ResponseWriter, r *http.Request) {
	s.serverHttp(w, r)
}

func (s *server) B(w http.ResponseWriter, r *http.Request) {
	s.serverHttp(w, r)
}

func (s *server) C(w http.ResponseWriter, r *http.Request) {
	s.serverHttp(w, r)
}

func okHttpServer() {
	s := &server{trace: openTrace("ok-http-server-service")}
	http.HandleFunc("/A", s.A)
	http.HandleFunc("/B", s.B)
	http.HandleFunc("/C", s.C)
	log.Error(http.ListenAndServe(":1234", nil))
}

func runServer() {
	go func() {
		okHttpServer()
	}()
	time.Sleep(time.Second)
}
