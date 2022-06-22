package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	for i := 1; i < 10; i++ {
		inter, err := net.InterfaceByIndex(i)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(inter.Name, inter.HardwareAddr.String())
	}

}
