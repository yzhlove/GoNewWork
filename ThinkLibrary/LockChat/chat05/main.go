package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

// http post 之multipart/form-data

func main() {

	go server(":8899")
	time.Sleep(time.Second * 1)
	data, err := httpPostMultipart("http://localhost:8899", map[string]string{"username": "yurisa", "passwd": "123456"})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("resp body:", string(data))

	time.Sleep(time.Second * 3)
}

func server(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("URL:%q", r.URL)
		log.Printf("Host:%q", r.Host)
		for k, v := range r.Header {
			log.Printf("Head: key:%s value:%v", k, v)
		}
		log.Println("Body:")
		data, _ := io.ReadAll(r.Body)
		log.Println(string(data))
		log.Println()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("request succeed."))
	})
	http.ListenAndServe(port, nil)
}

func httpPostMultipart(url string, params map[string]string) ([]byte, error) {

	body := bytes.NewBuffer([]byte{})
	port := multipart.NewWriter(body)

	if len(params) != 0 {
		for k, v := range params {
			if err := port.WriteField(k, v); err != nil {
				return nil, err
			}
		}
	}
	// 参数添加完了一定记得close
	port.Close()

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	cc := &http.Client{Timeout: time.Second * 5}
	resp, err := cc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
