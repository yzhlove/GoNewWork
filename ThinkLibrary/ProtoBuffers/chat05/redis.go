package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"runtime"
	"time"
)

var pool *redis.Pool

func Init(host string, index int) {
	pool = &redis.Pool{
		MaxIdle:     10 * runtime.NumCPU(),
		MaxActive:   50 * runtime.NumCPU(),
		Wait:        true,
		IdleTimeout: time.Minute,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host,
				redis.DialDatabase(index),
				redis.DialConnectTimeout(time.Second*2))
		},
	}

	for {
		if conn := pool.Get(); conn.Err() != nil {
			log.Print("redis get conn error:", conn.Err())
			conn.Close()
		} else {
			if _, err := conn.Do("PING"); err != nil {
				log.Print("redis ping error:", err)
			} else {
				conn.Close()
				break
			}
		}
		time.Sleep(time.Second)
	}
}

func Get() *redis.Pool {
	return pool
}
