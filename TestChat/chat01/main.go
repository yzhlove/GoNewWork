package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	var test actors
	test.add(func() error {
		fmt.Println("1 running...")
		time.Sleep(time.Second * 5)
		return errors.New("1 run stop")
	}, func(err error) {
		fmt.Println("1 err msg --> ", err)
	})

	test.add(func() error {
		fmt.Println("2 running...")
		time.Sleep(time.Second * 3)
		fmt.Println("2 exit...")
		return nil
	}, func(err error) {
		fmt.Println("2 err msg --> ", err)
	})

	test.add(func() error {
		fmt.Println("3 running...")
		time.Sleep(time.Second * 6)
		return errors.New("3 run stop")
	}, func(err error) {
		fmt.Println("3 err msg --> ", err)
	})

	fmt.Println(test.run())

	time.Sleep(time.Second * 2)
}

type actor struct {
	run func() error
	die func(err error)
}

type actors struct {
	data []actor
}

func (a *actors) add(run func() error, die func(err error)) {
	a.data = append(a.data, actor{run: run, die: die})
}

func (a *actors) run() error {
	if len(a.data) == 0 {
		return nil
	}

	erch := make(chan error, len(a.data))
	for _, k := range a.data {
		go func(k actor) {
			erch <- k.run()
		}(k)
	}

	err := <-erch
	for _, k := range a.data {
		k.die(err)
	}

	for i := 1; i < cap(erch); i++ {
		<-erch
	}
	return err
}
