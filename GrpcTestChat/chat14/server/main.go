package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"grpc-test-chat/chat14/echo"
	"hash/crc32"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	errTimeout = errors.New("handler request timeout")
)

const (
	maxQueue = 128
	maxGroup = 16
)

type (
	doReqMode int
	selfCtx   struct{}
)

const (
	Normal doReqMode = iota
	Rand
)

func main() {

	l, err := net.Listen("tcp", ":55566")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	sv := &service{serialQueueService: newQueue(ctx)}
	sv.run()

	serv := grpc.NewServer(grpc.UnaryInterceptor(sv.Intercept))
	echo.RegisterEchoServiceServer(serv, sv)

	log.Println("listing on ", l.Addr())
	if err := serv.Serve(l); err != nil {
		log.Panic(err)
	}
	cancel()

}

type Req struct {
	in     interface{}
	ctx    context.Context
	invoke grpc.UnaryHandler
	fn     func(resp Resp)
}

type Resp struct {
	out interface{}
	err error
}

type serialQueueService struct {
	chanQueue [maxGroup]chan *Req
	pool      *sync.Pool
	ctx       context.Context
	counter   uint32
}

func newQueue(ctx context.Context) *serialQueueService {
	return &serialQueueService{
		pool: &sync.Pool{
			New: func() interface{} {
				return &Req{}
			},
		},
		ctx: ctx,
	}
}

func (sq *serialQueueService) run() {
	for i := 0; i < maxGroup; i++ {
		sq.chanQueue[i] = make(chan *Req, maxQueue)
		go func(i int) {
			for {
				select {
				case <-sq.ctx.Done():
					return
				case req := <-sq.chanQueue[i]:
					req.ctx = context.WithValue(req.ctx, selfCtx{}, i+1)
					resp, err := req.invoke(req.ctx, req.in)
					req.fn(Resp{resp, err})
					sq.release(req)
				}
			}
		}(i)
	}
}

func (sq *serialQueueService) doReq(mode doReqMode, method string, req *Req) {
	var index int
	switch mode {
	case Normal:
		index = int(crc32.ChecksumIEEE([]byte(method)) % maxGroup)
	case Rand:
		index = int(sq.counter % maxGroup)
		atomic.AddUint32(&sq.counter, 1)
	}

	sq.chanQueue[index] <- req
}

func (sq *serialQueueService) acquire(ctx context.Context, in interface{}, invoke grpc.UnaryHandler, fn func(resp Resp)) *Req {
	req := sq.pool.Get().(*Req)
	req.ctx = ctx
	req.invoke = invoke
	req.in = in
	req.fn = fn
	return req
}

func (sq *serialQueueService) release(req *Req) {
	req.in = nil
	req.invoke = nil
	sq.pool.Put(req)

	return
}

type service struct {
	*serialQueueService
}

func (s *service) Echo(ctx context.Context, in *echo.Req) (*echo.Resp, error) {

	number, ok := ctx.Value(selfCtx{}).(int)
	if ok {
		log.Println("handle request form number => ", number)
	}

	var sum int
	for _, k := range strings.Split(in.Msg, "-") {
		v, err := strconv.Atoi(k)
		if err != nil {
			return nil, err
		}
		sum += v
	}

	return &echo.Resp{Code: 200, Msg: fmt.Sprintf("value:[%d]", sum)}, nil
}

func (s *service) Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, invoke grpc.UnaryHandler) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	respCh := make(chan Resp)
	request := s.acquire(ctx, req, invoke, func(resp Resp) { respCh <- resp })

	s.doReq(Rand, info.FullMethod, request)

	select {
	case <-ctx.Done():
		return nil, errTimeout
	case ret := <-respCh:
		return ret.out, ret.err
	}
}
