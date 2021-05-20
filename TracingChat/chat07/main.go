package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"io"
	"net/http"
	"strings"
	"time"
	"trace-chat/tracing"
)

func main() {

	stat := make(chan struct{})
	go beginService(stat)

	time.Sleep(time.Second)

	start()

	<-stat
}

func start() {

	t, c, err := tracing.NewOpentracing("chat07-test-trace-example")
	defer c.Close()
	_ = c
	_ = err

	span := t.StartSpan("first-span")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	timeFormatString(ctx)
	printShowString(ctx)

}

func timeFormatString(ctx context.Context) error {

	local := "http://localhost:1234/format"

	req, err := tracing.NewRequest(ctx, "time-format", local, tracing.H{"event": "time_format"})
	if err != nil {
		return err
	}

	resp, err := tracing.OkHttp(req)
	if err != nil {
		return err
	}

	fmt.Println("time-format result:", resp)
	return nil
}

func printShowString(ctx context.Context) error {

	local := "http://localhost:1234/show"

	req, err := tracing.NewRequest(ctx, "print-show", local, tracing.H{"event": "print_show"})
	if err != nil {
		return err
	}

	resp, err := tracing.OkHttp(req)
	if err != nil {
		return err
	}

	fmt.Println("print-show result:", resp)
	return nil
}

type server struct {
	trace opentracing.Tracer
	clos  io.Closer
}

func beginService(stat chan struct{}) {

	trace, clos, err := tracing.NewOpentracing("chat07-server-service")
	_ = clos
	if err != nil {
		panic(err)
	}
	s := &server{trace: trace}

	http.HandleFunc("/format", s.timeFormatService)
	http.HandleFunc("/show", s.printShowService)
	log.Error(http.ListenAndServe(":1234", nil))

	close(stat)
}

func (s *server) timeFormatService(w http.ResponseWriter, r *http.Request) {

	fmt.Println("time format server ....")

	span, err := tracing.ObtainSpan(s.trace, "time-format-server", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := opentracing.ContextWithSpan(r.Context(), span)
	newSpan, newContext := opentracing.StartSpanFromContext(ctx, "time-format-func-running")
	defer newSpan.Finish()

	str := fmt.Sprintf("hello-time-service:%s", r.FormValue("event"))
	newSpan.SetTag("input", "event")
	newSpan.SetTag("value", str)

	str = timeFormat(newContext, str)

	w.Write([]byte(str))
}

func timeFormat(ctx context.Context, event string) string {

	span, _ := opentracing.StartSpanFromContext(ctx, "time-format-func")
	defer span.Finish()

	s := strings.ToUpper(event)

	tracing.H{"event": "time-format-func", "value": s}.Iterator(func(k, v string) {
		span.SetTag(k, v)
		span.LogFields(log.String(k, v))
	})

	return s

}

func (s *server) printShowService(w http.ResponseWriter, r *http.Request) {

	fmt.Println("print show server ....")

	span, err := tracing.ObtainSpan(s.trace, "print-show-server", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer span.Finish()

	str := fmt.Sprintf("hello-print-service:%s", r.FormValue("event"))
	span.SetTag("input", "event")
	span.SetTag("value", str)

	w.Write([]byte(str))
}
