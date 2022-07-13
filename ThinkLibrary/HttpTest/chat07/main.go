package main

import (
	"bytes"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {

	go serverHandle()

	str := "0,1,2,3,4,5,6,7,8,9,10"
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9527", strings.NewReader(str))
	if err != nil {
		panic(err)
	}

	_, err = retryDo(req, time.Millisecond*500, 20, func(index int) time.Duration {
		return time.Millisecond * time.Duration(rand.Intn(5000)+1000)
	})

	if err != nil {
		panic(err)
	}

	time.Sleep(time.Hour)
}

func serverHandle() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("read request body error:%v", err)
			return
		}
		log.Printf("body len: %d", len(body))
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":9527", nil)
}

type retryFunc func(index int) time.Duration

func retryDo(req *http.Request, timeout time.Duration, retryCount int, rtfunc retryFunc) (*http.Response, error) {

	var bytesBody []byte
	var err error

	if req != nil && req.Body != nil {
		if bytesBody, err = copyReqBody(req.Body); err != nil {
			return nil, err
		}
		resetReqBody(req, bytesBody)
	}

	cc := &http.Client{Timeout: timeout}

	// 每次请求copy新的request
	copyRequest := func() *http.Request {
		r := req.Clone(req.Context())
		if r.Body != nil {
			resetReqBody(r, bytesBody)
		}
		return r
	}

	type reqRecord struct {
		resp  *http.Response
		err   error
		retry int
	}

	mulitpleCh := make(chan reqRecord)

	var wait = sync.WaitGroup{}
	var stopCh = make(chan struct{})

	go func() {
		wait.Wait()
		close(stopCh)
	}()

	for i := 0; i < retryCount; i++ {
		wait.Add(1)
		go func(i int) {
			time.Sleep(rtfunc(i) + time.Millisecond*100)

			defer wait.Done()
			resp, err := cc.Do(copyRequest())
			if err != nil {
				log.Printf("request error:%v", err)
			}
			if err == nil && resp.StatusCode < http.StatusInternalServerError {
				mulitpleCh <- reqRecord{
					resp:  resp,
					err:   err,
					retry: i,
				}
				return
			}

		}(i)
	}

	return nil, err
}

func copyReqBody(src io.ReadCloser) ([]byte, error) {
	bytes, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}
	defer src.Close()
	return bytes, nil
}

func resetReqBody(req *http.Request, data []byte) {
	req.Body = io.NopCloser(bytes.NewBuffer(data))
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewBuffer(data)), nil
	}
}
