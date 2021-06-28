package main

import (
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main() {

	cc, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	defer cc.Close()

}
