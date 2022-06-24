package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var ch = make(chan int, 1)
	var wg sync.WaitGroup
	var stop = make(chan struct{})

	wg.Add(4)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				return
			default:
				time.Sleep(time.Second * 5)
				value := rand.Intn(1000)
				log.Println("create value by:", value)
				ch <- value
			}
		}
	}()

	for i := 0; i < 3; i++ {
		go func(index int) {
			defer wg.Done()
			log.Println(index+1, " waiting receive value.")
			for {
				select {
				case <-stop:
					return
				case value := <-ch:
					log.Printf("[%d] get value:%d", index+1, value)
				}
			}
		}(i)
	}

	wg.Wait()

}
