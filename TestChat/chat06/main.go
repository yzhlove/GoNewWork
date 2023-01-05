package main

import (
	"fmt"
	"sync"
)

func main() {
	testswitchonce2()
}

func switchonce1(msg string, a, b chan<- string) {
	for i := 0; i < 2; i++ {
		select {
		case a <- msg:
			a = nil
		case b <- msg:
			b = nil
		}
	}
}

func testswitchonce1() {

	a := make(chan string)
	b := make(chan string)

	go switchonce1("hello world", a, b)

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case v := <-a:
				fmt.Println("value a -> ", v)
			case v := <-b:
				fmt.Println("value b -> ", v)
			}
		}()
	}
	wg.Wait()
}

func switchonce2(msg string, a, b chan<- string) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		a <- msg
	}()
	go func() {
		defer wg.Done()
		b <- msg
	}()
	wg.Wait()
}

func testswitchonce2() {

	a := make(chan string)
	b := make(chan string)

	go switchonce2("hello world", a, b)

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case v := <-a:
				fmt.Println("value a -> ", v)
			case v := <-b:
				fmt.Println("value b -> ", v)
			}
		}()
	}
	wg.Wait()
}
