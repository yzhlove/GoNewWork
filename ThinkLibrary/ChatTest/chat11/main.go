package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://raw.githubusercontent.com/yzhlove/GoDevWork/master/succeed.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if err := save(resp.Body); err != nil {
		log.Fatal(err)
	}

	log.Printf("ok .")

}

func save(reader io.Reader) error {

	buf := make([]byte, 4*1024)

	f, err := os.Create("apache-zookeeper.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = io.CopyBuffer(f, reader, buf)

	return err
}
