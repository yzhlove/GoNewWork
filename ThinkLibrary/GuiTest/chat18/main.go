package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	nets, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, inter := range nets {
		fmt.Printf("name:%s address->%s \n", inter.Name, inter.HardwareAddr.String())
		address, err := inter.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range address {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					fmt.Println(ipnet.IP.String())
					fmt.Println()
				}
			}
		}
	}

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		panic(err)
	}

	fmt.Println("----------------------------------------------------")

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {

			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}

		}
	}

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Println("name ----> ", name)

}
