package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/rand"
	"net"
	"time"
)

//daytime

const (
	localhost = "localhost:50051"
	port      = ":50051"
)

func main() {

	go server()
	if err := client(); err != nil {
		log.Print(err)
	}
}

func server() {
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	go handle(conn)
}

func handle(conn *net.UDPConn) {
	data := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Print(err)
			return
		}

		now := time.Now().Unix()
		log.Printf("[%s] read message:%s \n", time.Unix(now, 0).Format(time.RFC3339), string(data[:n]))
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(now))
		if _, err := conn.WriteToUDP(buf, addr); err != nil {
			log.Print(err)
		}
	}
}

func client() error {

	addr, err := net.ResolveUDPAddr("udp", localhost)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}

	buf := make([]byte, 8)
	for {
		if _, err := conn.Write(obtStr()); err != nil {
			return err
		}
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		tim := binary.BigEndian.Uint64(buf[:n])
		log.Printf("read time data: %s \n", time.Unix(int64(tim), 0).Format(time.RFC3339))
		time.Sleep(time.Second * 5)
	}
}

func obtStr() []byte {
	str := "abcdefghijklmnopqrstuvwxyz0192837465"
	var data bytes.Buffer
	for _, k := range rand.Perm(len(str)) {
		data.WriteString(string(str[k]))
	}
	return data.Bytes()
}
