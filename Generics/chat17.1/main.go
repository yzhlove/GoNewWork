package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

func main() {

	file, err := os.Open("/Users/yostar/workSpace/gowork/src/GoNewWork/Generics/chat20/main.go")
	if err != nil {
		panic(err)
	}

	stat, err := file.Stat()
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
		go handle(conn, file.Fd(), stat.Size())
	}
}

func handle(conn net.Conn, fd uintptr, size int64) {
	raw, err := conn.(*net.TCPConn).SyscallConn()
	if err != nil {
		fmt.Println("connection error", err)
		return
	}

	var sockFd uintptr
	if err = raw.Control(func(fd uintptr) {
		sockFd = fd
	}); err != nil {
		fmt.Println("connection error", err)
		return
	}

	var offset int64
	ws, err := syscall.Sendfile(int(sockFd), int(fd), &offset, int(size))
	if err != nil {
		fmt.Println("connection error", err)
		return
	}

	fmt.Println("send success bytes size:", ws)
}
