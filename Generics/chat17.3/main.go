package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	echoServer(&wg, ":8849")
	createdListen(&wg, ":8848", handleReadAndWrite)
	wg.Wait()

}

func handleReadAndWrite(conn net.Conn) {

	c, err := net.Dial("tcp", "localhost:8849")
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return
	}

	var totalSize int
	buf := make([]byte, 128)
	for {
		size, err := conn.Read(buf)
		fmt.Println("read ---> ", string(buf[:size]))
		if err != nil {
			fmt.Println("----------------- conn.read error:", err)
			break
		}
		if _, err = c.Write(buf[:size]); err != nil {
			fmt.Println("<handleReadAndWrite> write error: ", err)
			return
		} else {
			totalSize += size
			fmt.Println("<handleReadAndWrite> send bytes:", totalSize)
		}
	}
}

func handleEcho(conn net.Conn) {
	for {
		if _, err := io.Copy(conn, conn); err != nil {
			fmt.Println("<handleEcho> io.Copy error:", err)
		}
	}
}

func createdListen(wg *sync.WaitGroup, port string, handle func(conn net.Conn)) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}

	fmt.Println("net.Listen on ", port)
	go func() {
		defer wg.Done()
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("listen.Accept error:", err)
				continue
			}
			go handle(conn)
		}
	}()
}

func echoServer(wg *sync.WaitGroup, port string) {

	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}

	fmt.Println("net.Listen on ", port)
	var manager address
	var mutex sync.Mutex
	addFunc := func(at addr) {
		mutex.Lock()
		defer mutex.Unlock()
		manager.add(at)
	}

	handleFunc := func(conn net.Conn) {
		var buf = make([]byte, 64)
		for {
			size, err := conn.Read(buf)
			if err != nil {
				fmt.Println("<handleEcho>  error:", err)
				return
			}
			mutex.Lock()
			manager.iterator(func(at addr) {
				if at.ip == conn.RemoteAddr().String() {
					return
				}
				if _, err := at.conn.Write(buf[:size]); err != nil {
					fmt.Println("<handleEcho>  error:", at.ip, err)
				}
			})
			mutex.Unlock()
		}
	}

	go func() {
		defer wg.Done()
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("listen.Accept error:", err)
				return
			}
			fmt.Println("conn.RemoteAddr --> ", conn.RemoteAddr())
			addFunc(addr{conn.RemoteAddr().String(), conn})
			go handleFunc(conn)
		}
	}()

}

type addr struct {
	ip   string
	conn net.Conn
}

type address []addr

func (a *address) add(at addr) {
	for _, t := range *a {
		if t.ip == at.ip {
			return
		}
	}
	*a = append(*a, at)
}

func (a *address) iterator(callback func(at addr)) {
	for _, at := range *a {
		callback(at)
	}
}
