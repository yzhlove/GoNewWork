package main

import (
	"downloader/cmd"
	"downloader/down"
	"log"
)

func main() {

	//strUrl := "https://raw.githubusercontent.com/yzhlove/GoDevWork/master/succeed.jpg"

	c := cmd.ShowText()

	d := down.New(c.Concurrency)
	if err := d.Download(c.DownURL, c.Filename); err != nil {
		log.Fatal(err)
	}
	log.Printf("down file {%s} succ.", c.Filename)
}
