package main

import (
	"encoding/base64"
	"fmt"
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

func Benchmark_Db(b *testing.B) {

	key := "test-set-db-key"
	value := base64.StdEncoding.EncodeToString([]byte("hello world"))

	redispool := setPool("127.0.0.1:6379")
	keydbpool := setPool("127.0.0.1:6380")

	b.Run("test-redis-write", func(b *testing.B) {

		conn := redispool.Get()
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

		conn := keydbpool.Get()
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

	b.Run("test-redis-read", func(b *testing.B) {
		conn := redispool.Get()
		defer conn.Close()

		b.ResetTimer()
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			retV, err := redis.String(conn.Do("GET", key))
			if err != nil {
				b.Error(err)
				return
			}
			if retV != value {
				b.Error(fmt.Sprintf("read error"))
			}
		}

		b.StopTimer()
		b.ReportAllocs()
	})

	b.Run("test-keydb-read", func(b *testing.B) {
		conn := keydbpool.Get()
		defer conn.Close()

		b.ResetTimer()
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			retV, err := redis.String(conn.Do("GET", key))
			if err != nil {
				b.Error(err)
				return
			}
			if retV != value {
				b.Error(fmt.Sprintf("read error"))
			}
		}

		b.StopTimer()
		b.ReportAllocs()
	})

}

/*
Benchmark_Db
Benchmark_Db/test-redis-write
Benchmark_Db/test-redis-write-12         	     936	   1225344 ns/op	      64 B/op	       3 allocs/op
Benchmark_Db/test-keydb-write
Benchmark_Db/test-keydb-write-12         	    1293	    922294 ns/op	      64 B/op	       3 allocs/op
Benchmark_Db/test-redis-read
Benchmark_Db/test-redis-read-12          	    1423	    835452 ns/op	      88 B/op	       5 allocs/op
Benchmark_Db/test-keydb-read
Benchmark_Db/test-keydb-read-12          	    1443	    935608 ns/op	      88 B/op	       5 allocs/op
PASS
*/

func Benchmark_Db_Parallel(b *testing.B) {

	key := "test-set-db-key-parallel"
	value := base64.StdEncoding.EncodeToString([]byte("hello world"))

	redispool := setPool("127.0.0.1:6380")
	//keydbpool := setPool("127.0.0.1:6380")

	b.ResetTimer()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		conn := redispool.Get()
		defer conn.Close()
		for pb.Next() {
			if _, err := conn.Do("SET", key, value); err != nil {
				b.Error(err)
				return
			}
		}
	})
	b.StopTimer()
	b.ReportAllocs()

}

/*
Benchmark_Db_Parallel
Benchmark_Db_Parallel-12    	    5330	    222745 ns/op	      87 B/op	       3 allocs/op
PASS

Benchmark_Db_Parallel
Benchmark_Db_Parallel-12    	    6796	    166016 ns/op	      83 B/op	       3 allocs/op
PASS
*/
