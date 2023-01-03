package main

import (
	"errors"
	"fmt"
	"github.com/oklog/run"
	"time"
)

func main() {
	var test run.Group

	test.Add(func() error {
		fmt.Println("1 running...")
		time.Sleep(time.Second * 5)
		return errors.New("1 run stop")
	}, func(err error) {
		fmt.Println("1 err msg --> ", err)
	})

	test.Add(func() error {
		fmt.Println("2 running...")
		time.Sleep(time.Second * 3)
		fmt.Println("2 exit...")
		return nil
	}, func(err error) {
		fmt.Println("2 err msg --> ", err)
	})

	test.Add(func() error {
		fmt.Println("3 running...")
		time.Sleep(time.Second * 6)
		return errors.New("3 run stop")
	}, func(err error) {
		fmt.Println("3 err msg --> ", err)
	})

	fmt.Println("test.Run -> ", test.Run())
	time.Sleep(time.Second * 2)
}