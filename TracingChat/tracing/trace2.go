package tracing

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
	"strings"
)

type H map[string]string

func (h H) Encoder() string {
	vars := make([]string, 0, len(h)>>1)
	for k, v := range h {
		vars = append(vars, k+"="+v)
	}
	return "?" + strings.Join(vars, "&")
}

func (h H) Iterator(f func(k, v string)) {
	for k, v := range h {
		f(k, v)
	}
}

func NewOpentracing(service string) (opentracing.Tracer, io.Closer, error) {

	cfg := &config.Configuration{
		ServiceName: service,
		Sampler:     &config.SamplerConfig{Type: "const", Param: 1}, //全量采集
		Reporter:    &config.ReporterConfig{LogSpans: true},
	}

	trace, clos, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Errorf("new trace error:%v", err))
	}

	opentracing.SetGlobalTracer(trace)
	return trace, clos, err
}

func OkHttp(req *http.Request) (string, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func NewRequest(ctx context.Context, service, url string, vars H) (*http.Request, error) {

	span, _ := opentracing.StartSpanFromContext(ctx, "client:"+service)
	defer span.Finish()

	httpUrl := url + vars.Encoder()
	fmt.Println("encoder url->", httpUrl, vars)

	req, err := http.NewRequest(http.MethodGet, httpUrl, nil)
	if err != nil {
		return nil, err
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, httpUrl)
	ext.HTTPMethod.Set(span, req.Method)

	err = span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	vars.Iterator(func(k, v string) {
		span.SetTag(k, v)
		span.LogFields(log.String(k, v))
	})

	return req, err
}

func ObtainSpan(trace opentracing.Tracer, service string, req *http.Request) (opentracing.Span, error) {

	ctx, err := trace.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

	span := trace.StartSpan("server:"+service, ext.RPCServerOption(ctx))
	if err != nil {
		span = trace.StartSpan("server-span")
	}
	defer span.Finish()

	H{"service": service, "host": req.Host, "url": req.URL.String(), "method": req.Method}.
		Iterator(func(k, v string) {
			span.SetTag(k, v)
			span.LogFields(log.String(k, v))
		})
	return span, nil
}
