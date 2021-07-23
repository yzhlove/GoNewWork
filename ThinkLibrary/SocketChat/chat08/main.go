package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

//标准库多播

func main() {

	addr, err := net.ResolveUDPAddr("udp", "224.0.0.255:9999")
	if err != nil {
		log.Fatal(err)
	}

	//en, err := net.InterfaceByName("en1")
	//if err != nil {
	//	log.Fatal(err)
	//}

	cc, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listener address: %s \n", cc.LocalAddr())

	buf := make([]byte, 1024)
	for {
		n, remote, err := cc.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("remote addr: %s remote data: %s \n", remote, string(buf[:n]))

		str := strings.ToUpper(fmt.Sprintf("[%s->%s]:%s", addr, remote, string(buf[:n])))
		log.Printf("write remote {%s} data: {%s} \n", remote, str)
		if _, err := cc.WriteToUDP([]byte(str), remote); err != nil {
			log.Fatal(err)
		}
	}

}
