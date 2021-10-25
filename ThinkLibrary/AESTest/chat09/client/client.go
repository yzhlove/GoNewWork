package main

import (
	"log"
	"net"
)

type handlefunc func(conn net.Conn, c *client) error

type client struct {
	address string
	hfunc   handlefunc
	conn    net.Conn
}

func newClient(address string, hfunc handlefunc) *client {
	c := &client{address: address, hfunc: hfunc}
	return c
}

func (c *client) start() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", c.address)
	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}

	c.conn = conn
	return nil
}

func (c *client) run() {
	if err := c.hfunc(c.conn, c); err != nil {
		log.Println("run func error:", err)
		c.stop()
	}
}

func (c *client) stop() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func handler(conn net.Conn, c *client) error {

	return nil
}

func main() {

}
