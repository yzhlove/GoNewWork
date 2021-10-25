package main

import (
	"context"
	"log"
	"net"
	"think-library/AESTest/chat09_mode/pack"
)

type connect struct {
	conn  net.Conn
	msgCh chan []byte
}

func newConnect(conn net.Conn) *connect {
	return &connect{conn: conn, msgCh: make(chan []byte, 1)}
}

func (c *connect) startReader() {
	parse := pack.Parser{Conn: c.conn}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		if packet, ok := parse.Next(); ok {
			if f, ok := MsgCenter[packet.Id()]; ok {
				c.msgCh <- f(ctx, packet)
			} else {
				log.Println("find msgId error:", packet.Id())
			}
		} else {
			break
		}
	}
	log.Println("parse error:", parse.Err())
}

func (c *connect) startWrite() {
	for msg := range c.msgCh {

	}
}
