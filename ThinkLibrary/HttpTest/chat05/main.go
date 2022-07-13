package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	go serverHandle()

	str := "0,1,2,3,4,5,6,7,8,9,10"
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9527", strings.NewReader(str))
	if err != nil {
		panic(err)
	}

	resp, err := retryDo(req, 10, time.Millisecond*500, func(i int) time.Duration {
		return time.Millisecond * 1000
	})
	if err != nil {
		panic(err)
	}

	resp = resp
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

type backoffStrategy func(i int) time.Duration

func retryDo(req *http.Request, maxRetries int, timeout time.Duration, backoff backoffStrategy) (*http.Response, error) {

	var originBody []byte
	var err error

	if req != nil && req.Body != nil {
		if originBody, err = copyBody(req.Body); err != nil {
			return nil, err
		}
		resetBody(req, originBody)
	}

	attempt := maxRetries
	if attempt <= 0 {
		attempt = 1
	}

	cc := &http.Client{Timeout: timeout}
	var resp *http.Response

	for i := 0; i < attempt; i++ {
		if resp, err = cc.Do(req); err != nil {
			log.Printf("request error: %v", err)
		}
		// 针对500以上的错误码进行重试
		if err == nil && resp.StatusCode < 500 {
			return resp, nil
		}
		//if resp != nil {
		//	resp.Body.Close()
		//}
		//if resp != nil && resp.Body != nil {
		resetBody(req, originBody)
		//}
		time.Sleep(backoff(i) + time.Microsecond)
	}

	return resp, req.Context().Err()
}

func copyBody(src io.ReadCloser) ([]byte, error) {
	b, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}
	src.Close()
	return b, nil
}

func resetBody(request *http.Request, originBody []byte) {
	request.Body = io.NopCloser(bytes.NewBuffer(originBody))
	request.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewBuffer(originBody)), nil
	}
}
