package main

import (
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

const localhost = "localhost"
const port = "50051"
const protocol = "udp"

func main() {
	go server()
	client()
}

func server() {

	addr, err := net.ResolveUDPAddr(protocol, ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP(protocol, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024*4)

	for {
		n, resAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("recv message:%s \n", string(buf[0:n]))

		msg := strings.ToUpper(string(buf[0:n]))

		if _, err = conn.WriteToUDP([]byte(msg), resAddr); err != nil {
			log.Fatal(err)
		}
	}

}

func client() {

	conn, err := net.Dial(protocol, localhost+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	gener := func() string {
		str := "abcdefghijklmnopqrstuvwxyz"
		var tmp strings.Builder
		for _, k := range rand.Perm(len(str) >> 1) {
			tmp.WriteString(string(str[k]))
		}
		return tmp.String()
	}

	buf := make([]byte, 4*1024)

	for {
		if _, err := conn.Write([]byte(gener())); err != nil {
			log.Fatal(err)
		}

		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("read server msg:%s \n", string(buf[0:n]))
		time.Sleep(time.Second * 5)
	}
}
