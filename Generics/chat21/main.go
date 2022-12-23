package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTicker(time.Second * 3)
	ch := make(chan struct{})
	var index int
	go func() {
		<-ch
		fmt.Println("waiting over.")
		time.Sleep(time.Second)
		index = 0
		timer.Reset(time.Second * 5)
	}()
	go func() {
		for {
			<-timer.C
			if index >= 5 {
				timer.Stop()
				ch <- struct{}{}
			}
			fmt.Println("tick.", index)
			index++
		}
	}()

	time.Sleep(time.Hour)

}
