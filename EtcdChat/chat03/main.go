package main

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
)

func main() {
	cc, err := clientv3.New(clientv3.Config{Endpoints: []string{"localhost:2379"}})
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	kv := clientv3.NewKV(cc)

	resp, err := kv.Put(context.Background(), "/servers/1", "192.168.1.123")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("resp --> ", resp)

	resp, err = kv.Put(context.Background(), "/servers/2", "192.168.1.234")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("resp --> ", resp)

	resp, err = kv.Put(context.Background(), "/servers/3", "192.168.2.456")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("resp --> ", resp)

	res, err := kv.Get(context.Background(), "/servers", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < int(res.Count); i++ {
		log.Println("kv -> ", res.Kvs[i].String())
	}

}
