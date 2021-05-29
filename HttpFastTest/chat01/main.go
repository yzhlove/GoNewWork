package main

import (
	"http-fast/pressure"
	"log"
	"net/http"
	"time"
)

func main() {

	go pressure.NewHttpServer()

	fast := NewHttpFast(pressure.Number, "http://localhost:1234/")
	fast.Run()

}

func NewHttpFast(count int, url string) *HttpFast {
	return &HttpFast{
		Count:    count,
		Host:     url,
		reqChan:  make(chan *ReqTask, 128),
		bufChan:  make(chan *ReqTask, 256),
		respChan: make(chan *RespTask, 128),
		stat:     make(chan struct{}),
	}
}

type HttpFast struct {
	Count    int
	Host     string
	reqChan  chan *ReqTask
	bufChan  chan *ReqTask
	respChan chan *RespTask
	stat     chan struct{}
	ts       time.Time
}

type ReqTask struct {
	Id  int
	Req *http.Request
}

type RespTask struct {
	Id  int
	Err error
}

func (fast *HttpFast) Run() {
	go fast.show()
	go fast.sendRequest()
	go fast.deliverTask()
	fast.exec()
	<-fast.stat
}

func (fast *HttpFast) exec() {
	defer close(fast.reqChan)
	fast.ts = time.Now()
	start := time.Now()
	for i := 0; i < fast.Count; i++ {
		if req, err := http.NewRequest(http.MethodGet, fast.Host, nil); err != nil {
			log.Print("create request error:", err)
		} else {
			if i != 0 && i%100 == 0 {
				end := time.Now()
				log.Print("create http request by:", i, end.Sub(start))
				start = end
			}
			fast.reqChan <- &ReqTask{Id: i + 1, Req: req}
		}
	}

}

func (fast *HttpFast) deliverTask() {
	defer close(fast.bufChan)
	for task := range fast.reqChan {
		fast.bufChan <- task
	}
}

func (fast *HttpFast) sendRequest() {
	defer close(fast.respChan)
	client := http.Client{Timeout: time.Second * 2}
	for task := range fast.bufChan {
		_, err := client.Do(task.Req)
		fast.respChan <- &RespTask{Id: task.Id, Err: err}
	}
}

func (fast *HttpFast) show() {
	var succeed, failed int
	for resp := range fast.respChan {
		if err := resp.Err; err != nil {
			failed++
		} else {
			succeed++
		}
	}
	log.Printf("- Request URL{%s} Count{%d} Succeed:%d Failed:%d Stat:%.3f%% Time:%v -",
		fast.Host, fast.Count, succeed, failed, float32(succeed)/float32(fast.Count), time.Now().Sub(fast.ts))
	close(fast.stat)
}
