package model

import "sync"

type Application struct {
	sync.RWMutex
	appid         string
	instances     map[string]*Instance
	lastTimestamp int64
}

