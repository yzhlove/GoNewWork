package main

import (
	"encoding/base64"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func Test_Redis(t *testing.T) {

	conn := pool.Get()
	defer conn.Close()

	value := "hello world"

	_, err := conn.Do("SET", "test-redis", value)
	if err != nil {
		t.Error(err)
		return
	}

	ret, err := redis.String(conn.Do("GET", "test-redis"))
	if err != nil {
		t.Error(err)
		return
	}

	if ret != value {
		t.Error("test failed")
	} else {
		t.Log("Ok.")
	}
}

func Test_Keydb(t *testing.T) {

	Host = "127.0.0.1:6380"

	conn := pool.Get()
	defer conn.Close()

	value := "hello world"

	_, err := conn.Do("SET", "test-keydb", value)
	if err != nil {
		t.Error(err)
		return
	}

	ret, err := redis.String(conn.Do("GET", "test-keydb"))
	if err != nil {
		t.Error(err)
		return
	}

	if ret != value {
		t.Error("test failed")
	} else {
		t.Log("Ok.")
	}
}

func Benchmark_Db_Write(b *testing.B) {

	key := "test-set-db-key"
	value := base64.StdEncoding.EncodeToString([]byte("hello world"))

	b.Run("test-redis-write", func(b *testing.B) {

		pool := setPool("127.0.0.1:6379")
		conn := pool.Get()
		defer conn.Close()

		b.ResetTimer()
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			if _, err := conn.Do("SET", key, value); err != nil {
				b.Error(err)
				return
			}
		}

		b.StopTimer()
		b.ReportAllocs()
	})

	b.Run("test-keydb-write", func(b *testing.B) {

		pool := setPool("127.0.0.1:6380")
		conn := pool.Get()
		defer conn.Close()

		b.ResetTimer()
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			if _, err := conn.Do("SET", key, value); err != nil {
				b.Error(err)
				return
			}
		}

		b.StopTimer()
		b.ReportAllocs()
	})

}
