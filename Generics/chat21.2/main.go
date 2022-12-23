package main

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {

	secondTick := time.NewTicker(time.Second * 5)
	minuteTick := time.NewTicker(time.Second * 10)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	log := zap.NewExample()

check:
	for {
		if err := sync(); err != nil {
			log.Error("sync error", zap.Error(err))
			select {
			case <-ctx.Done():
				log.Debug("done", zap.String("position", "one"))
				return
			case <-secondTick.C:
				log.Debug("secondTick")
				continue check
			}
		}
		log.Info("one succeed exit")
		secondTick.Stop()
		break
	}

	log.Info("two process join ...")
	minuteTick.Reset(time.Second * 10)
	for {
		select {
		case <-ctx.Done():
			log.Debug("done2", zap.String("position", "two"))
			return
		case <-minuteTick.C:
			log.Debug("minuteTick start")
			if err := sync(); err != nil {
				log.Error("sync2 error", zap.Error(err))

				minuteTick.Stop()
				secondTick.Reset(time.Second * 5)
				goto check
			}
			log.Debug("minuteTick stop")
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
