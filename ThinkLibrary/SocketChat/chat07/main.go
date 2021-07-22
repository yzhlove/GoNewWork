package main

import (
	"log"
	"net"
	"time"
)

func main() {

	ip := net.ParseIP("224.0.0.255")

	cc, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: ip, Port: 9999})
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()
	log.Print("local  addr: ", cc.LocalAddr())
	log.Print("remote addr: ", cc.RemoteAddr())

	buf := make([]byte, 1500) //mtu size
	for {
		_, err := cc.Write([]byte("hello world!"))
		if err != nil {
			log.Fatal(err)
		}
		n, err := cc.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("read from:{%s} data: %s \n", cc.RemoteAddr(), string(buf[:n]))
		time.Sleep(time.Second)
	}

}
