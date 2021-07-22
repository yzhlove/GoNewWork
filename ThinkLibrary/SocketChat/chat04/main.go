package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const (
	localhost = "localhost"
	port      = 9999
)

func main() {
	go server()
	//client()
	client2()
}

func server() {

	cc, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP(localhost), Port: port})
	if err != nil {
		log.Fatal(err)
	}

	log.Print("listener to address: ", cc.LocalAddr().String())

	buf := make([]byte, 1024)
	for {
		n, addr, err := cc.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		str := strings.ToUpper(string(buf[:n])) + "\n"
		if _, err = cc.WriteToUDP([]byte(str), addr); err != nil {
			log.Fatal(err)
		}
	}
}

func client() {
	src := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dst := &net.UDPAddr{IP: net.ParseIP(localhost), Port: port}

	cc, err := net.DialUDP("udp", src, dst)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	log.Print("remote addr: ", cc.LocalAddr().String(), "->", cc.RemoteAddr())

	if _, err := cc.Write([]byte("hello world")); err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, err := cc.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("read remote data : ", string(buf[:n]))

	time.Sleep(time.Second)
}

func client2() {

	src := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dst := &net.UDPAddr{IP: net.ParseIP(localhost), Port: port}

	cc, err := net.DialUDP("udp", src, dst)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	go func() {
		scan := bufio.NewScanner(cc)
		for scan.Scan() {
			log.Println("scan data: ", scan.Text())
		}
	}()
	log.Println("remote addr: ", cc.RemoteAddr())
	for {
		buf := make([]byte, 20)
		os.Stdin.Read(buf)
		cc.Write(buf)
	}
}
