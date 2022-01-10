package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// 自旋锁，不支持可重入

type spinLocker uint32

func (spin *spinLocker) Lock() {
	for atomic.CompareAndSwapUint32((*uint32)(spin), 0, 1) {
		runtime.Gosched()
	}
}

func (spin *spinLocker) Unlock() {
	atomic.StoreUint32((*uint32)(spin), 0)
}

func NewSpinLocker() sync.Locker {
	var spin spinLocker
	return &spin
}
