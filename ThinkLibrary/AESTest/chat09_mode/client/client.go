package main

import (
	"net"
	"think-library/AESTest/chat09_mode/conf"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", conf.HostName)
	if err != nil {

	}
}

