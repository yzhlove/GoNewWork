package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

func main() {

	tick := time.NewTicker(time.Minute * 10)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	file, err := os.Create("httpMonitor.log")
	if err != nil {
		log.Fatalf("创建监听文件失败:%v", err)
	}

	sw := &syncWrite{Writer: file}
	var count = 0

	defer func() {
		file.Sync()
		file.Close()
	}()

	log.Printf("进程执行中...\n")

	for {
		count++
		select {
		case <-tick.C:
			log.Printf("[%s][%d]定时器执行中。。。\n", time.Now().Format(timeLayout), count)

			ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
			run(ctx, sw)
			file.Sync()

		case <-sigCh:
			log.Printf("进程退出.\n")
			return
		}
	}

}

const (
	ipAddress  = "http://47.100.202.166:8008/manager/meta/serverlist.html"
	urlAddress = "http://nova-sh-dev-oss.oss-accelerate.aliyuncs.com/meta/serverlist.html"
)

var (
	address = []string{ipAddress, urlAddress}
)

func run(ctx context.Context, w io.Writer) {
	var wg sync.WaitGroup
	for _, addr := range address {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				log.Printf("[%s] 超时退出!", addr)
				break
			default:
				writeHttpResp(addr, w)
			}
		}(addr)
	}
	wg.Wait()
}

func writeHttpResp(url string, w io.Writer) {
	data, err := getHttpResponse(url)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("[%s][%s] 请求发生错误:[%v]\n", time.Now().Format(timeLayout), url, err)))
		return
	}
	if err := testParse(data); err != nil {
		w.Write([]byte(fmt.Sprintf("[%s][%s][%s] 请求返回的数据:[%v]\n", time.Now().Format(timeLayout), "FALSE", url, string(data))))
	} else {
		w.Write([]byte(fmt.Sprintf("[%s][%s][%s] 请求返回的数据:[%v]\n", time.Now().Format(timeLayout), "TRUE", url, string(data))))
	}
}

func getHttpResponse(url string) ([]byte, error) {
	cc := &http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := cc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func testParse(data []byte) error {
	mgr := &Manager{}
	return json.Unmarshal(data, mgr)
}

type Manager struct {
	Version int    `json:"V"`           // 资源版本
	Agent   []Gate `json:"A,omitempty"` // 区服入口
}

type Gate struct {
	Name   string `json:"N"`           // 名称
	Addr   string `json:"A,omitempty"` // 地址端口
	Status int    `json:"S,omitempty"` // 状态
	Zone   int    `json:"Z,omitempty"` // 区
	Count  int    `json:"-"`           // 多gate计数
}

type syncWrite struct {
	io.Writer
	sync.Mutex
}

func (sw *syncWrite) Write(p []byte) (n int, err error) {
	sw.Lock()
	defer sw.Unlock()

	return sw.Writer.Write(p)
}
