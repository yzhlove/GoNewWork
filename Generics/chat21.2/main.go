package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {

	secondTick := time.NewTicker(time.Second * 5)
	minuteTick := time.NewTicker(time.Second * 10)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

check:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done.")
			return
		case <-secondTick.C:
			fmt.Println("secondTick.")
			if err := sync(); err != nil {
				fmt.Println("sync error:", err)
				continue check
			}
		}
		secondTick.Stop()
		break
	}

	minuteTick.Reset(time.Second * 10)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done2.")
			return
		case <-minuteTick.C:
			fmt.Println("minuteTick.")
			if err := sync(); err != nil {
				fmt.Println("sync error:", err)

				secondTick.Reset(time.Second * 5)
				goto check
			}
		}
	}

}

func sync() error {
	resp, err := http.Get("http://127.0.0.1:7887")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("status code error")
	}
	return nil
}
