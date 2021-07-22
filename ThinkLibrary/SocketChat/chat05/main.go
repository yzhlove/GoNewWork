package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

//等价的客户端和服务端

func main() {
	echo()
}

func read(cc *net.UDPConn) {
	buf := make([]byte, 1024)
	for {
		n, err := cc.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("{%s} read data: %s \n\n", cc.LocalAddr(), string(buf[:n]))
	}
}

func echo() {

	addr1 := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981}
	addr2 := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9982}

	go func() {
		cc, err := net.ListenUDP("udp", addr1)
		if err != nil {
			log.Fatal(err)
		}
		go read(cc)
		var count int
		for {
			str := fmt.Sprintf("[%d] ping: {%s} -> {%s}", count, cc.LocalAddr(), addr2)
			if _, err := cc.WriteToUDP([]byte(str), addr2); err != nil {
				log.Fatal(err)
			}
			count += 2
			time.Sleep(time.Second)
		}
	}()

	go func() {
		cc, err := net.ListenUDP("udp", addr2)
		if err != nil {
			log.Fatal(err)
		}
		go read(cc)
		var count = 1
		for {
			str := fmt.Sprintf("[%d] pong: {%s} -> {%s}", count, cc.LocalAddr(), addr1)
			if _, err := cc.WriteToUDP([]byte(str), addr1); err != nil {
				log.Fatal(err)
			}
			count += 2
			time.Sleep(time.Second)
		}
	}()

	buf := make([]byte, 1)
	os.Stdin.Read(buf)
}
