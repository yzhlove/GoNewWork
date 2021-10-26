package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"think-library/AESTest/chat09_mode/conf"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", conf.HostName)
	if err != nil {
		toErr("ResolveTCPAddr", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		toErr("ListenTCP", err)
	}

	log.Println("listener to port:", conf.HostName)

	for {
		conn, err := listener.Accept()
		if err != nil {
			toErr("Accept", err)
		}
		go handler(conn)
	}

}

func handler(conn net.Conn) {
	defer conn.Close()

	ctx := NewContext()

	go func() {
		for msg := range ctx.msgCh {
			if _, err := conn.Write(msg); err != nil {
				toErr("Write", err)
			}
		}
	}()

	p := parser{conn: conn}
	for {
		data, err := p.read()
		if err != nil {
			toErr("parser", err)
		}
		req := ctx.req(data)
		if req != nil {
			handle := Get(req.Id())
			if err := ctx.send(handle(ctx, req)); err != nil {
				toErr("send", err)
			}
		}
	}

}

type parser struct {
	conn net.Conn
}

func (p *parser) read() ([]byte, error) {
	var size uint32
	if err := binary.Read(p.conn, binary.LittleEndian, &size); err != nil {
		return nil, err
	}

	if size == 0 {
		return nil, errors.New("read conn size not is zero")
	}

	buf := make([]byte, size)
	//if err := binary.Read(p.conn, binary.LittleEndian, &data); err != nil {
	//	return nil, err
	//}

	if _, err := io.ReadFull(p.conn, buf); err != nil {
		return nil, err
	}

	return buf, nil
}

func toErr(prefix string, err error) {
	panic(fmt.Sprintf("[%s] panic:%v", prefix, err))
}
