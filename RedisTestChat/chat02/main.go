package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"redis-chat/redisc"
)

func main() {

	redisc.Init("127.0.0.1:6379", 2)

	p := redisc.Get()
	test1(p)
	test2(p)

}

func test1(p *redis.Pool) {
	//values := make(Values, 0, 4)
	//values = append(values, )
	var values Values
	values[0] = V{Left: 0, Right: 100}

	data, err := values.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}

	cc := p.Get()
	defer cc.Close()

	if _, err = cc.Do("SET", "RedisTypeValues", data); err != nil {
		panic(err)
	}

	fmt.Println("write ok.")
}

func test2(p *redis.Pool) {

	v := V{Left: 0, Right: 100}

	data, err := v.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}

	cc := p.Get()
	defer cc.Close()

	if _, err = cc.Do("SET", "RedisTypeSignValue", data); err != nil {
		panic(err)
	}

	fmt.Println("write ok.")

}
