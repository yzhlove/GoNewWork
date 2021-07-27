package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type handleFunc func(ctx context.Context, conn net.Conn)

type HandlerInter interface {
	Handler(ctx context.Context, conn net.Conn)
	Close() error
}

type config struct {
	address    string
	maxConnect uint32
	timeout    time.Duration
}

func listenerAndServerWithSignal(cfg *config, handler HandlerInter) error {
	closeCh := make(chan struct{})
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigCh
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeCh <- struct{}{}
		}
	}()

	listener, err := net.Listen("tcp", cfg.address)
	if err != nil {
		return err
	}

	//TODO:
	log.Printf("bind:%s start listener...", cfg.address)

	listenerAndServer(listener, handler, closeCh)
	return nil
}

func listenerAndServer(listener net.Listener, handler HandlerInter, closeChan <-chan struct{}) {

	go func() {
		<-closeChan
		listener.Close()
		handler.Close()
	}()

	defer func() {
		listener.Close()
		handler.Close()
	}()

	ctx := context.Background()
	var wait sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}

		log.Printf("accept link...")

		wait.Add(1)
		go func() {
			defer func() {
				wait.Done()
			}()
			handler.Handler(ctx, conn)
		}()
	}
	wait.Wait()
}

type boolean uint32

func (b *boolean) Get() bool {
	return atomic.LoadUint32((*uint32)(b)) != 0
}

func (b *boolean) Set(v bool) {
	if v {
		atomic.StoreUint32((*uint32)(b), 1)
	} else {
		atomic.StoreUint32((*uint32)(b), 0)
	}
}

type Wait struct {
	wg sync.WaitGroup
}

func (w *Wait) Add(delta int) {
	w.wg.Add(delta)
}

func (w *Wait) Done() {
	w.wg.Done()
}

func (w *Wait) Wait() {
	w.wg.Wait()
}

func (w *Wait) WithWaitTimeout(timeout time.Duration) bool {
	c := make(chan bool, 1)
	go func() {
		defer close(c)
		w.wg.Wait()
		c <- true
	}()
	select {
	case <-c:
		return false
	case <-time.After(timeout):
		return true
	}
}

type EchoHandler struct {
	activeConn sync.Map
	closing    boolean
}

func MakeEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

type EchoClient struct {
	Conn    net.Conn
	Waiting Wait
}

func (c *EchoClient) Close() error {
	c.Waiting.WithWaitTimeout(time.Second * 10)
	c.Conn.Close()
	return nil
}

func (h *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	if h.closing.Get() {
		conn.Close()
	}

	cc := &EchoClient{Conn: conn}
	h.activeConn.Store(cc, struct{}{})

	reader := bufio.NewReader(conn)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("connect close ...")
				h.activeConn.Delete(cc)
			} else {
				log.Printf("reader data error: %v \n", err)
			}
			return
		}

		cc.Waiting.Add(1)
		conn.Write([]byte(data))
		cc.Waiting.Done()
	}
}

func (h *EchoHandler) Close() error {
	log.Printf("handler shutting down...")
	h.closing.Set(true)
	h.activeConn.Range(func(key, value interface{}) bool {
		cc := key.(*EchoClient)
		cc.Close()
		return true
	})
	return nil
}

func main() {

}
