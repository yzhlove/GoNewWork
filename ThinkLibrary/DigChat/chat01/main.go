package main

import (
	"bytes"
	"go.uber.org/dig"
	"log"
	"os"
)

func main() {

	type Config struct {
		Prefix string
	}

	c := dig.New()

	err := c.Provide(func() (*Config, error) {
		return &Config{Prefix: "[provide]"}, nil
	})
	toErr(err)

	err = c.Provide(func(cfg *Config) *log.Logger {
		return log.New(os.Stdout, cfg.Prefix, 0)
	})
	toErr(err)

	err = c.Invoke(func(l *log.Logger) {
		l.Println("call invoke func")
	})
	toErr(err)

	buf := &bytes.Buffer{}
	err = dig.Visualize(c, buf)
	toErr(err)

	f, err := os.Create("dig.dot")
	if err != nil {
		toErr(err)
	}
	f.Write(buf.Bytes())
	f.Close()

}

func toErr(err error) {
	if err != nil {
		panic(err)
	}
}
