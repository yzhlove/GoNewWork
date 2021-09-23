package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"think-library/AESTest/chat04_tcp/config"
	"time"
)

func main() {
	ch := make(chan struct{})
	server()
	<-ch
}

func server() error {
	l, err := net.Listen("tcp", config.HOST)
	if err != nil {
		return err
	}

	conn, err := l.Accept()
	if err != nil {
		return err
	}
	fmt.Println("accept addr -> ", conn.RemoteAddr())
	go handle(conn)

	return nil
}

func handle(conn net.Conn) {

	var rpm int
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		rpm++

		str, err := reader.ReadString('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatal(err)
			}
			return
		}

		log.Println("reader:", str)

		writer.WriteString(fmt.Sprintf("%d:%s\n", rpm, strings.ToUpper(str)))
		if err := writer.Flush(); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 2)
	}

}
