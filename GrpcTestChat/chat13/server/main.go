package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"grpc-test-chat/chat13/echo"
	"hash/crc32"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	maxQueue  = 128
	groupSize = 32
)

func main() {
	l, err := net.Listen("tcp", ":55566")
	if err != nil {
		panic(err)
	}

	svc := NewService()
	svc.run()
	sever := grpc.NewServer()
	echo.RegisterEchoServiceServer(sever, svc)
	fmt.Println("listing on the :", l.Addr())

	if err := sever.Serve(l); err != nil {
		log.Fatal(err)
	}
}

type callbackFunc func(resp Resp)
type serialHandleFunc func(req *Req) Resp

type Req struct {
	req *echo.Req
	fn  callbackFunc
}

type Resp struct {
	resp *echo.Resp
	err  error
}

type service struct {
	serialQueue [groupSize]chan *Req
	pool        *sync.Pool
	ctx         context.Context
	cancel      context.CancelFunc
	handle      serialHandleFunc
}

func NewService() *service {
	s := &service{}
	for k := range s.serialQueue {
		s.serialQueue[k] = make(chan *Req, maxQueue)
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.handle = handleFunc
	s.pool = &sync.Pool{
		New: func() interface{} {
			return &Req{req: &echo.Req{}}
		},
	}
	return s
}

func (s *service) run() {
	for k := range s.serialQueue {
		go func(k int) {
			for {
				select {
				case <-s.ctx.Done():
					return
				case req, ok := <-s.serialQueue[k]:
					if ok {
						resp := s.handle(req)
						req.fn(resp)
					}
				}
			}
		}(k)
	}
}

func (s *service) stop() error {
	s.cancel()
	return nil
}

func (s *service) take(r *echo.Req, fn callbackFunc) *Req {
	ret := s.pool.Get().(*Req)
	if ret.req == nil {
		ret.req = &echo.Req{Msg: r.Msg}
	} else {
		ret.req.Msg = r.Msg
	}
	ret.fn = fn
	return ret
}

func (s *service) release(in *Req) {
	in.req.Msg = ""
	in.fn = nil
	s.pool.Put(in)
}

func (s *service) doReq(in *Req) {
	index := crc32.ChecksumIEEE([]byte(in.req.Msg)) % groupSize
	s.serialQueue[index] <- in
}

func (s *service) Echo(ctx context.Context, req *echo.Req) (*echo.Resp, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := make(chan Resp)

	in := s.take(req, func(resp Resp) { res <- resp })
	defer s.release(in)

	s.doReq(in)

	select {
	case <-ctx.Done():
		return nil, errors.New("[handle] request handle timeout ")
	case resp := <-res:
		return resp.resp, resp.err
	}
}

func handleFunc(in *Req) Resp {
	var sum int
	for _, value := range strings.Split(in.req.Msg, "-") {
		r, err := strconv.Atoi(value)
		if err != nil {
			return Resp{err: err}
		}
		sum += r
	}
	return Resp{resp: &echo.Resp{Code: 200, Msg: fmt.Sprintf("res:[%d]", sum)}}
}
