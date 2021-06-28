package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"log"
)

func main() {
	cc, err := clientv3.New(clientv3.Config{Endpoints: []string{"localhost:2379"}})
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	kv := clientv3.NewKV(cc)
	resp, err := kv.Put(context.Background(), "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
