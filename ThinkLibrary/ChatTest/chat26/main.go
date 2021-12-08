package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	st := New()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		st.set("test", rand.Perm(10)...)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Microsecond * 1)
		st.add("test", 100)
	}()

	wg.Wait()

}

type stat struct {
	sync.RWMutex
	data map[string][]int
}

func New() *stat {
	return &stat{data: make(map[string][]int)}
}

func (s *stat) set(a string, b ...int) {
	s.RLock()
	defer s.RUnlock()
	fmt.Println("read lock")
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	s.data[a] = b
}

func (s *stat) add(a string, b int) {
	s.Lock()
	defer s.Unlock()
	fmt.Println("locker")
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	s.data[a] = append(s.data[a], b)
}
