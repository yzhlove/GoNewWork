package main

import (
	"golang.org/x/net/ipv4"
	"log"
	"net"
)

//UDP多播

func main() {

	//得到一块网卡
	en1, err := net.InterfaceByName("en1")
	if err != nil {
		log.Fatal(err)
	}

	group := net.IPv4(224, 0, 0, 255)

	cc, err := net.ListenPacket("udp4", "0.0.0.0:9999")
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	pack := ipv4.NewPacketConn(cc)
	if err := pack.JoinGroup(en1, &net.UDPAddr{IP: group}); err != nil {
		log.Fatal(err)
	}

	if err := pack.SetControlMessage(ipv4.FlagDst, true); err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1500) //mtu size
	for {
		n, cm, src, err := pack.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}
		//如果是组播
		if cm.Dst.IsMulticast() {
			//检查是不是同一个分组
			if cm.Dst.Equal(group) {
				log.Printf("receive data: %s from: {%s}\n", buf[:n], src.String())
				if n, err := pack.WriteTo([]byte("abcde hello world !!!"), cm, src); err != nil {
					log.Fatal(err)
				} else {
					log.Printf("send {%s} data length: %d \n", src, n)
				}
			} else {
				log.Printf("unknown group: %s \n", src.String())
				continue
			}
		}
	}

}
