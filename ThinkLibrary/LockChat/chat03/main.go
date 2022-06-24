package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	factoryCh := make(chan int)

	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Second * 10)
			value := rand.Intn(1000)
			select {
			case factoryCh <- value:
				log.Println("factory value:", value)
			default:
				log.Println("factory skip:", value)
			}
		}
	}()

	go func() {
		defer wg.Done()
		index := 1
		for {
			index++
			if index%2 == 0 {
				select {
				case value := <-factoryCh:
					log.Println("get value --> ", value)
				}
			}
		}
	}()

	wg.Wait()
}
