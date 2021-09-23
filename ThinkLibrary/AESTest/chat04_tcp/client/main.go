package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"think-library/AESTest/chat04_tcp/config"
)

func main() {
	log.Fatal(client())
}

func client() error {
	conn, err := net.Dial("tcp", config.HOST)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	var count int

	for {
		count++
		if _, err := writer.WriteString(fmt.Sprintf("%d:%s\n", count, "hello world")); err != nil {
			log.Fatal(err)
		}
		if err := writer.Flush(); err != nil {
			log.Fatal(err)
		}
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		log.Println("reader string->", str)
	}

	return nil
}
