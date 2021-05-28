package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"redis-chat/redisc"
)

func main() {

	redisc.Init("127.0.0.1:6379", 1)

	//test2()

	test3()

}

func test1() {

	pool := redisc.Get()

	_, err := redisc.Save(pool.Get(), func(conn redis.Conn) error {
		s := redisc.NewSelf(conn)
		s.Send("SET", "key_a", "12345").
			Send("SET", "key_b", "2345").
			Send("SET", "key_c", "1232145321545")
		return s.Err()
	})

	if err != nil {
		log.Print("redis pipeline error:", err)
	} else {
		log.Print("redis succeed ok.")
	}

}

func test2() {

	pool := redisc.Get()

	_, err := redisc.Save(pool.Get(), func(conn redis.Conn) error {
		s := redisc.NewSelf(conn)
		s.Send("INCR", "A", "B", "C")
		return s.Err()
	})

	if err != nil {
		log.Print("redis pipeline error:", err)
	} else {
		log.Print("redis succeed ok.")
	}

}

func test3() {

	pool := redisc.Get()

	reply, err := redisc.Save(pool.Get(), func(conn redis.Conn) error {
		s := redisc.NewSelf(conn)
		s.Send("SET", "a", 300)
		s.Send("LPOP", "a")
		return s.Err()
	})

	if err != nil {
		log.Print("redis pipeline error:", err)
	} else {
		log.Print("redis succeed ok.")
	}

	for _, r := range reply.([]interface{}) {
		if err, ok := r.(redis.Error); ok {
			log.Print("reply error->", err)
		}
	}

}
