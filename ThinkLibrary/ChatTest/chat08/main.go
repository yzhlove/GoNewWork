package main

import (
	"runtime"
	"sync/atomic"
)

func main() {

}

type spin uint32

func (sp *spin) Lock() {
	for atomic.CompareAndSwapUint32((*uint32)(sp), 0, 1) {
		runtime.Gosched()
	}
}

func (sp *spin) Unlock() {
	atomic.StoreUint32((*uint32)(sp), 0)
}

type self struct {
	size int
	cnt  int
	fs   []float32
	l    spin
	//l sync.RWMutex
}

func new(sz int) *self {
	return &self{size: sz, fs: make([]float32, 0, sz)}
}

func (s *self) put(a float32) {
	s.l.Lock()
	defer s.l.Unlock()
	if s.size == s.cnt {
		runtime.Gosched()
		return
	}
	s.fs = append(s.fs, a)
	s.cnt++
}

func (s *self) pop() float32 {
	//s.l.RLock()
	//defer s.l.RUnlock()
	s.l.Lock()
	defer s.l.Unlock()
	if s.cnt == 0 {
		runtime.Gosched()
		return 0
	}
	s.cnt--
	return s.fs[s.cnt]
}

func (s *self) GetInt(max, min int32) int32 {
	if a := s.pop(); a != 0 {
		return int32(float32(max-min+1)*a) + min
	}
	return 0
}
