package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"grpc-test-chat/chat11/echo"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type task struct {
	in  *echo.Req
	out chan *echo.Resp
}

type service struct {
	reqCh chan *task
}

func New() *service {
	return &service{
		reqCh: make(chan *task, 512),
	}
}

func (s *service) run() {
	go func() {
		for tk := range s.reqCh {
			strs := strings.Split(tk.in.Msg, "-")
			var count int
			for _, str := range strs {
				ret, err := strconv.Atoi(str)
				if err != nil {
					tk.out <- nil
					continue
				}
				count += ret
			}
			tk.out <- &echo.Resp{Msg: fmt.Sprintf("values:[%v] count:[%d]", strs, count), Code: 200}
		}
	}()
}

func (s *service) Do(tk *task) {
	s.reqCh <- tk
}

func (s *service) Echo(ctx context.Context, in *echo.Req) (*echo.Resp, error) {

	respCh := make(chan *echo.Resp, 1)
	timer := time.NewTimer(time.Second * 5)

	s.Do(&task{in: in, out: respCh})

	select {
	case ret, ok := <-respCh:
		timer.Stop()
		if ok {
			if ret != nil {
				return ret, nil
			}
			return nil, errors.New("return resp nil")
		}
	case <-timer.C:
		return nil, errors.New("timeout")
	}

	return nil, errors.New("handle request error")
}

func main() {
	l, err := net.Listen("tcp", ":55566")
	if err != nil {
		panic(err)
	}

	svc := New()
	svc.run()
	sever := grpc.NewServer()
	echo.RegisterEchoServiceServer(sever, svc)
	fmt.Println("listing on the :", l.Addr())

	if err := sever.Serve(l); err != nil {
		log.Fatal(err)
	}
}
