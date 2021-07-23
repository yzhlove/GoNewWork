package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

//UDP广播

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go server(&wg)

	go client1(&wg)
	//go client2(&wg)

	wg.Wait()

}

func server(wg *sync.WaitGroup) {

	defer wg.Done()

	cc, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 9981})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[server][listener on address: {%s}] \n", cc.LocalAddr())

	buf := make([]byte, 1024)
	for {
		n, remote, err := cc.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("[server]receive remote [%s] data: %s \n", remote, string(buf[:n]))

		msg := fmt.Sprintf("server sendto message: [%s->%s]{%s}", cc.LocalAddr(), remote, string(buf[:n]))

		if _, err := cc.WriteToUDP([]byte(msg), remote); err != nil {
			log.Fatal(err)
		}
	}
}

func server2(wg *sync.WaitGroup) {

	defer wg.Done()

	cc, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 9981})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[server][listener on address: {%s}] \n", cc.LocalAddr())

	buf := make([]byte, 1024)
	for {
		n, remote, err := cc.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("[server2]receive remote [%s] data: %s \n", remote, string(buf[:n]))

		//msg := fmt.Sprintf("server sendto message: [%s->%s]{%s}", cc.LocalAddr(), remote, string(buf[:n]))
		//
		//if _, err := cc.WriteToUDP([]byte(msg), remote); err != nil {
		//	log.Fatal(err)
		//}
	}
}

func client1(wg *sync.WaitGroup) {

	defer wg.Done()

	src := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dst := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981}

	cc, err := net.ListenUDP("udp", src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[client1][listener on address: {%s}] \n", cc.LocalAddr())

	buf := make([]byte, 1024)
	for {

		if _, err := cc.WriteToUDP([]byte("hello world"), dst); err != nil {
			log.Fatal(err)
		}

		n, remote, err := cc.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[client1]receive remote [%s] data: %s \n", remote, string(buf[:n]))

		time.Sleep(time.Second * 10)
	}
}

func client2(wg *sync.WaitGroup) {

	defer wg.Done()

	src := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dst := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981}

	cc, err := net.ListenUDP("udp", src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[client2][listener on address: {%s}] \n", cc.LocalAddr())

	buf := make([]byte, 1024)
	for {

		if _, err := cc.WriteToUDP([]byte("hello world"), dst); err != nil {
			log.Fatal(err)
		}

		n, remote, err := cc.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[client2]receive remote [%s] data: %s \n", remote, string(buf[:n]))

		time.Sleep(time.Second * 10)
	}
}
