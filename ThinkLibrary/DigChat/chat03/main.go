package main

import (
	"bytes"
	"go.uber.org/dig"
	"log"
	"os"
)

func main() {

	type DNS struct {
		Address string
	}

	type DNSService struct {
		dig.Out
		ADNS *DNS `name:"a_service"`
		BDNS *DNS `name:"b_service"`
	}

	c := dig.New()

	if err := c.Provide(func() (DNSService, error) {
		return DNSService{
			ADNS: &DNS{Address: "127.0.0.1:1087"},
			BDNS: &DNS{Address: "192.168.2.11:1087"}}, nil
	}); err != nil {
		panic(err)
	}

	type DataBaseInfo struct {
		dig.In
		DNS_A *DNS `name:"a_service"`
		DNS_B *DNS `name:"b_service"`
	}

	if err := c.Invoke(func(db DataBaseInfo) {
		log.Println(db.DNS_A)
		log.Println(db.DNS_B)
	}); err != nil {
		panic(err)
	}

	withDot(c)

}

func withDot(c *dig.Container) {
	var buf = &bytes.Buffer{}
	if err := dig.Visualize(c, buf); err != nil {
		panic(err)
	}
	f, err := os.Create("dig.dot")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(buf.Bytes())
}
