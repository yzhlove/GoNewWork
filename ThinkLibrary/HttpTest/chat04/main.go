package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
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
	str := "0123456789"
	buf := bytes.NewReader([]byte(str))

	req, err := http.NewRequest(http.MethodPost, "http://localhost:9527", buf)
	if err != nil {
		return err
	}

	cc := &http.Client{Timeout: time.Millisecond * 500}

	for {
		if _, err := cc.Do(req); err != nil {
			log.Printf("request err:%v", err)
		}
		buf.Seek(0, 0)
		time.Sleep(time.Second * 3)
	}

}
