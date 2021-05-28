package redisc

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
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host, redis.DialDatabase(index), redis.DialConnectTimeout(2*time.Second))
		},
	}
	checkConnection()
}

func checkConnection() {
	var conn redis.Conn
	rt := time.NewTicker(2 * time.Second)
	defer rt.Stop()
	for {
		if conn != nil {
			if _, err := conn.Do("PING"); err != nil {
				log.Print("redis ping error:", err)
			} else {
				return
			}
		} else {
			if conn = pool.Get(); conn.Err() != nil {
				log.Print("redis connect error:", conn.Err())
				conn.Close()
				conn = nil
			}
		}
		<-rt.C
	}
}

func Get() *redis.Pool {
	return pool
}

func Save(conn redis.Conn, fn func(conn redis.Conn) error) (interface{}, error) {

	if err := conn.Send("MULTI"); err != nil {
		return nil, err
	}
	if err := fn(conn); err != nil {
		return nil, err
	}
	return conn.Do("EXEC")
}

type self struct {
	err  error
	conn redis.Conn
}

func NewSelf(conn redis.Conn) *self {
	return &self{conn: conn}
}

func (s *self) Send(c string, args ...interface{}) *self {
	if s.err == nil {
		s.err = s.conn.Send(c, args...)
	}
	return s
}

func (s *self) Do(c string, args ...interface{}) *self {
	if s.err == nil {
		_, s.err = s.conn.Do(c, args...)
	}
	return s
}

func (s *self) Err() error {
	return s.err
}
