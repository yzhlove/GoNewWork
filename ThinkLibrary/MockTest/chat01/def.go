package main

//go:generate mockgen -source=def.go -package=main -destination=mock_def.go

import "log"

type DeferDataBase interface {
	Get(key string) (int, error)
}

func GetFromDataBase(db DeferDataBase, key string) int {
	res, err := db.Get(key)
	if err != nil {
		log.Printf("db search error:%v", err)
		return -1
	}
	return res
}
