package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

//UDP池化

func main() {
	stop := make(chan struct{})
	go server(stop)



	time.Sleep(time.Second * 5)
	close(stop)
}



func start(wg *sync.WaitGroup, address string, stop chan struct{}) {

	defer wg.Done()

	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal(err)
	}

	cc, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listener to [%s] \n", cc.LocalAddr())
	var count, failed int
	buf := make([]byte, 1500)

OUT:
	for {
		select {
		case <-stop:
			break OUT
		default:
			n, remote, err := cc.ReadFromUDP(buf)
			if err != nil {
				break OUT
			}
			if !check(remote.String(), count, string(buf[:n])) {
				failed++
			}
			count++
		}
	}
	log.Printf("[%s] info stat: count [%d] failed [%d] \n", cc.LocalAddr(), count, failed)
}

func check(remote string, count int, str string) bool {
	s := strings.Split(str, ":")
	code := md5.Sum([]byte(fmt.Sprintf("%s:%d", remote, count)))
	return s[len(s)-1] == fmt.Sprintf("%x", code)
}

func server(stop chan struct{}) {

	address := []string{":8800", ":8801", ":8802"}
	var wg sync.WaitGroup
	for _, addr := range address {
		wg.Add(1)
		go start(&wg, addr, stop)
	}
	wg.Wait()

}
