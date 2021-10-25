package main

import (
	"log"
	"net"
	"os"
	"sync"
)

type handlefunc func(conn net.Conn, s *server)

type server struct {
	address  string
	listener net.Listener
	sync.WaitGroup
	hfunc  handlefunc
	errCh  chan error
	stopCh chan struct{}
}

func newServer(address string, hfunc handlefunc) *server {
	s := &server{address: address, hfunc: hfunc}
	s.errCh = make(chan error, 1)
	s.stopCh = make(chan struct{})
	return s
}

func (s *server) start() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", s.address)
	if err != nil {
		return err
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}
	s.listener = l
	return nil
}

func (s *server) stop() {
	<-s.stopCh
	os.Exit(0)
}

func (s *server) run() {
	go func() {
		for err := range s.errCh {
			log.Println("handle func error:", err)
		}
	}()
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			break
		}
		s.Add(1)
		s.hfunc(conn, s)
	}
	s.Wait()

	s.stopCh <- struct{}{}
}

func handler(conn net.Conn, s *server) {
	defer s.Done()

}

func main() {

}
