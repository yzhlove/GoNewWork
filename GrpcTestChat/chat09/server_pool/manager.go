package main

import (
	"context"
	"google.golang.org/grpc"
)

type Req struct {
	method  string
	ctx     context.Context
	request interface{}
	handle  grpc.UnaryHandler
}

type Resp struct {
	resp interface{}
	err  error
}

type handleManager struct {
	reqCh  chan Req
	respCh chan Resp
}

func NewHandleManager() *handleManager {
	return &handleManager{
		reqCh:  make(chan Req, 128),
		respCh: make(chan Resp, 128),
	}
}

func (h *handleManager) DoReq(r Req) {
	h.reqCh <- r
}

func (h *handleManager) DoResp() (interface{}, error) {
	r := <-h.respCh
	return r.resp, r.err
}

func (h *handleManager) run() {
	go func() {
		for req := range h.reqCh {
			resp, err := req.handle(req.ctx, req.request)
			h.respCh <- Resp{resp: resp, err: err}
		}
	}()
}
