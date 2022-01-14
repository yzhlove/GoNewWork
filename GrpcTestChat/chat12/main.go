package main

import (
	"context"
	"fmt"
	"time"
)

// test context levels

func main() {

	ctx1, cancel := context.WithCancel(context.Background())

	ctx2, _ := context.WithCancel(ctx1)
	go func() {
		select {
		case <-ctx2.Done():
			fmt.Println("ctx2 exit.")
		}
	}()

	ctx6, _ := context.WithCancel(ctx2)
	go func() {
		select {
		case <-ctx6.Done():
			fmt.Println("ctx6 exit.")
		}
	}()

	ctx3, cancel3 := context.WithCancel(ctx1)
	go func() {
		select {
		case <-ctx3.Done():
			fmt.Println("ctx3 exit")
		}
	}()

	ctx4, _ := context.WithCancel(ctx3)
	go func() {
		select {
		case <-ctx4.Done():
			fmt.Println("ctx4 exit")
		}
	}()

	ctx5, cancel5 := context.WithCancel(ctx3)
	go func() {
		select {
		case <-ctx5.Done():
			fmt.Println("ctx5 exit")
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("5 stop ing ...")
	cancel5()

	time.Sleep(time.Second)
	fmt.Println("3 stop ing ...")
	cancel3()

	time.Sleep(time.Second)
	fmt.Println("1 stop ing ...")
	cancel()

	time.Sleep(time.Second)
	fmt.Println("Ok.")
}
