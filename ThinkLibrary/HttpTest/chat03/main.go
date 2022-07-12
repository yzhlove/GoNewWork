package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	go serverHandle()
	if err := httpDo(); err != nil {
		panic(err)
	}
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

func httpDo() error {
	str := "1234567890"
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9527", strings.NewReader(str))
	if err != nil {
		return err
	}
	cc := &http.Client{Timeout: time.Millisecond * 500}

	for {
		if _, err := cc.Do(req); err != nil {
			log.Printf("request error: %v", err)
		}
		time.Sleep(time.Second * 3)
	}
}
