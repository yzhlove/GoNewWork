package main

import (
	"log"
	"time"
)

func main() {
	go testTick3()
	time.Sleep(time.Second * 10)
}

func testTick() {

	rt := time.NewTicker(time.Second)
	select {
	case <-rt.C:
		log.Print("time tick event.")
	}

}

func testTick2() {
	select {
	case <-time.NewTicker(time.Second).C:
		log.Print("time tick2 event.")
	}
}

func testTick3() {
	rt := time.NewTicker(time.Second)
	for {
		select {
		case <-rt.C:
			log.Print("time.tick event.")
		}
	}
}
