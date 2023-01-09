package main

import (
	"github.com/gomodule/redigo/redis"
	"runtime"
	"time"
)

var pool *redis.Pool

var (
	Host = "127.0.0.1:6379"
)

func init() {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				Host,
				redis.DialDatabase(1),
				redis.DialConnectTimeout(time.Second*5),
				redis.DialReadTimeout(time.Second*5),
				redis.DialWriteTimeout(time.Second*5))
		},
		MaxIdle:     10 * runtime.NumCPU(),
		MaxActive:   50 * runtime.NumCPU(),
		IdleTimeout: 60 * time.Second,
		Wait:        true,
	}
}

func setPool(host string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				host,
				redis.DialDatabase(1),
				redis.DialConnectTimeout(time.Second*5),
				redis.DialReadTimeout(time.Second*5),
				redis.DialWriteTimeout(time.Second*5))
		},
		MaxIdle:     10 * runtime.NumCPU(),
		MaxActive:   50 * runtime.NumCPU(),
		IdleTimeout: 60 * time.Second,
		Wait:        true,
	}
}
