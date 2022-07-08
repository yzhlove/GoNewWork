package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	tm := newTimer(60)

	for i := 0; i < 10; i++ {
		ti := i
		tm.add(i, func() {
			log.Println("count i value:", ti)
		})
	}

	time.Sleep(time.Second * 30)
	tm.stop()
}

type Timer struct {
	*time.Ticker
	sync.RWMutex
	sync.WaitGroup
	sync.Once
	maxScale int
	scale    int
	tasks    map[int][]func()
	stopCh   chan struct{}
}

func newTimer(scale int) *Timer {
	rt := &Timer{maxScale: scale, tasks: make(map[int][]func(), scale), stopCh: make(chan struct{})}
	rt.run()
	return rt
}

func (t *Timer) add(delay int, task func()) {
	t.Lock()
	defer t.Unlock()

	index := delay % t.maxScale
	t.tasks[index] = append(t.tasks[index], task)
}

func (t *Timer) run() {
	t.Ticker = time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-t.Ticker.C:
				t.task()
			case <-t.stopCh:
				log.Println("timer stop.")
				return
			}
		}
	}()
}

func (t *Timer) task() {
	t.Lock()
	defer t.Unlock()

	if ret, ok := t.tasks[t.scale]; ok {
		for _, fn := range ret {
			t.Add(1)
			go func(fn func()) {
				defer t.Done()
				fn()
			}(fn)
		}
	}
	t.Wait()
	log.Printf("[%d] run task ok.", t.scale)
	t.scale++
	if t.scale%t.maxScale == 0 {
		t.scale = 0
	}
}

func (t *Timer) stop() {
	t.Once.Do(func() {
		close(t.stopCh)
	})
}
