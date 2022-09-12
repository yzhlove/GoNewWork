package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	ctx := NewContext()
	request1 := ctx.Get()
	request1.name = "One"
	ctx.httpDo(request1)

	for _, url := range []string{
		"http://www.baidu.com",
		"http://www.sina.com",
		"http://www.pixes.com",
	} {
		request1.put(url)
	}

	ctx.Put(request1)

	_ = ctx.Get()

	go func() {
		for _, url := range []string{
			"http://www.abc.com",
			"http://www.youku.com",
			"http://www.tudou.com",
		} {
			request1.put(url)
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)
		ctx.stop()
	}()
	ctx.wait()
}

type req struct {
	name    string
	request chan *http.Request
}

func (r *req) put(url string) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	r.request <- request
}

type context struct {
	*sync.Pool
	sync.WaitGroup
	stopCh chan struct{}
}

func (ctx *context) Get() *req {
	return ctx.Pool.Get().(*req)
}

func (ctx *context) Put(r *req) {
	ctx.Pool.Put(r)
}

func NewContext() *context {
	return &context{
		stopCh: make(chan struct{}),
		Pool: &sync.Pool{
			New: func() interface{} {
				return &req{
					request: make(chan *http.Request, 32),
				}
			},
		},
	}
}

func (ctx *context) httpDo(treq *req) {
	ctx.Add(1)
	go func() {
		defer ctx.Done()
		for {
			select {
			case ret := <-treq.request:
				resp, err := http.DefaultClient.Do(ret)
				if err != nil {
					fmt.Printf("[ERRO] [%s] %v\n", treq.name, err)
					break
				}
				_, err = io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("[ERRO] [%s] %v", treq.name, err)
					break
				}
				fmt.Printf("[INFO] [%s] [%s] \n", ret.URL, treq.name)
				resp.Body.Close()
			case <-ctx.stopCh:
				return
			}
		}
	}()
}

func (ctx *context) stop() {
	close(ctx.stopCh)
}

func (ctx *context) wait() {
	ctx.Wait()
}
