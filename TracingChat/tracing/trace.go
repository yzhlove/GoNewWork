package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

func NewTrace(service string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler:     &config.SamplerConfig{Type: "const", Param: 1},
		Reporter:    &config.ReporterConfig{LogSpans: true},
	}
	return cfg.NewTracer(config.Logger(jaeger.StdLogger))
}
