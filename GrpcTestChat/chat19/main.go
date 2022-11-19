package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gomodule/redigo/redis"
	"math/rand"
)

func main() {

	setData(createMeta())
	getData()
}

func createMeta() *Meta {
	meta := &Meta{}

	for i := 0; i < 3; i++ {
		resp := &Resp{}
		for k := 0; k < 10; k++ {
			resp.Items = append(resp.Items, &Resp_Item{
				Tid: uint32(rand.Intn(100) + 10),
				Qty: int32(rand.Intn(1000) - rand.Intn(2000)),
			})
		}
		meta.extra = append(meta.extra, resp)
	}
	meta.Sid = "test save proto message bytes"
	return meta
}

func setData(meta *Meta) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(1))
	if err != nil {
		panic(err)
	}

	if _, err = conn.Do("SET", "Test_Proto_Message", meta); err != nil {
		panic(err)
	}
	fmt.Println("Ok.")

}

func getData() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(1))
	if err != nil {
		panic(err)
	}

	reply, err := c.Do("GET", "Test_Proto_Message")
	if err != nil {
		panic(err)
	}

	meta := &Meta{}
	if _, err = redis.Scan([]any{reply}, meta); err != nil {
		panic(err)
	}

	fmt.Println(meta.Sid)
	for _, t := range meta.extra {
		fmt.Printf("%+v \n", t)
	}
}

//go:generate msgp -tests=false -io=false

type Meta struct {
	extra []*Resp `msg:"-"`
	Sid   string
	Bytes [][]byte
}

func (m Meta) RedisArg() any {

	for _, resp := range m.extra {
		bytes, err := proto.Marshal(resp)
		if err != nil {
			panic(err)
		}
		m.Bytes = append(m.Bytes, bytes)
	}

	data, err := m.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}
	return data
}

func (m *Meta) RedisScan(src any) error {

	if _, err := m.UnmarshalMsg(src.([]byte)); err != nil {
		panic(err)
	}

	for _, bytes := range m.Bytes {
		var resp = &Resp{}
		if err := proto.Unmarshal(bytes, resp); err != nil {
			panic(err)
		}
		m.extra = append(m.extra, resp)
	}
	return nil
}
