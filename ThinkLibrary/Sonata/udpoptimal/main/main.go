package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
	"upd-chat/cnc"
)

//UDP池化

func main() {
	stop := make(chan struct{})
	go server(stop)
	now := time.Now()
	client()
	log.Println("close client ... ", time.Now().Sub(now).Milliseconds(), "ms")
	close(stop)
	time.Sleep(time.Second * 10)
	log.Println("Ok.")
}

//1360ms
//45478ms

func client() {
	address := []string{"localhost:8800", "localhost:8801", "localhost:8802"}
	var wg sync.WaitGroup
	for _, addr := range address {
		wg.Add(1)
		go send(&wg, addr)
	}
	wg.Wait()
}

func send(wg *sync.WaitGroup, addr string) {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		if err := cnc.Report(addr, encode(i)); err != nil {
			log.Fatal(err)
		}
		//if err := cnc.Report2(addr, encode(addr, i)); err != nil {
		//	log.Fatal(err)
		//}
	}
}

func encode(index int) string {
	s := fmt.Sprintf("%d", index)
	ss := md5.Sum([]byte(s))
	return fmt.Sprintf("%s:%x", s, ss)
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
			n, _, err := cc.ReadFromUDP(buf)
			if err != nil {
				break OUT
			}
			if !check(count, string(buf[:n])) {
				failed++
			}
			count++
		}
	}
	log.Printf("[%s] info stat: count [%d] failed [%d] \n", cc.LocalAddr(), count, failed)
}

func check(count int, str string) bool {
	s := strings.Split(str, ":")
	code := md5.Sum([]byte(fmt.Sprintf("%d", count)))
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
