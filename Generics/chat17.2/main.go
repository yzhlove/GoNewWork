package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	file, err := os.Open("/Users/yostar/workSpace/gowork/src/GoNewWork/Generics/chat20/main.go")
	if err != nil {
		panic(err)
	}

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept:", err)
		}
		go handle(conn, file)
	}
}

func handle(conn net.Conn, reader io.Reader) {
	ws, err := io.Copy(conn, reader)
	if err != nil {
		fmt.Println("handle error:", err)
		return
	}
	fmt.Println("send size: ", ws)
}
