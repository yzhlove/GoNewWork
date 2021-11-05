package main

import (
	"bytes"
	"go.uber.org/dig"
	"log"
	"os"
)

func main() {

	type DNS struct {
		Addr string
	}

	c := dig.New()

	c.Provide(func() (*DNS, error) {
		return &DNS{Addr: "127.0.0.1:1087"}, nil
	}, dig.Name("a_service"))

	c.Provide(func() (*DNS, error) {
		return &DNS{Addr: "10.111.222.1:8080"}, nil
	}, dig.Name("b_service"))

	type LoadConfig struct {
		dig.In
		A_DNS *DNS `name:"a_service"`
		B_DNS *DNS `name:"b_service"`
	}

	c.Invoke(func(c *LoadConfig) {
		log.Println(c.A_DNS.Addr)
		log.Println(c.B_DNS.Addr)
	})

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
