package main

import (
	"log"
	"net"
	"os"
	"syscall"
)

func main() {

	file, err := os.Open("/Users/yostar/workSpace/gowork/src/GoNewWork/Generics/chat06/main.go")
	if err != nil {
		panic(err)
	}
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	l, err := net.Listen("tcp", ":8848")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept error: %v \n", err)
		}
		go handle(conn, file.Fd(), stat.Size())
	}

}

func handle(conn net.Conn, fd uintptr, size int64) {
	file, err := conn.(*net.TCPConn).File()
	if err != nil {
		log.Printf("net.Conn assest error: %v\n", err)
		return
	}

	var offset int64
	// Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
	// outfd 可以是文件也可以是网络，
	// intfd 必须是文件
	// offset 发送数据的起始位置
	// count 发送数据的大小
	wsize, err := syscall.Sendfile(int(file.Fd()), int(fd), &offset, int(size))
	if err != nil {
		log.Printf("sendfile error: %v \n", err)
		return
	}
	log.Printf("sendfile size is %d \n", wsize)
}
