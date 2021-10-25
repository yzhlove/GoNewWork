package server

import (
	"log"
	"net"
	"sync"
	"think-library/AESTest/chat07/conf"
)

type server struct {
	wg sync.WaitGroup
}

func (s *server) Listen() error {
	l, err := net.Listen("tcp", conf.Hostname)
	if err != nil {
		return err
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print("[server] accept error:", err)
			break
		}
		s.wg.Add(1)
		go s.handle(conn)
	}
	s.wg.Wait()
	return nil
}

func (s *server) handle(conn net.Conn) {
	defer s.wg.Done()

}
