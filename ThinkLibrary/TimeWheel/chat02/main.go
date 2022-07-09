package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	tr := newTimer(60)
	tr.add(130, func() {
		log.Println("✈︎ 130 running...")
	})
	tr.add(60, func() {
		log.Println("✈︎ 60 running...")
	})
	tr.add(30, func() {
		log.Println("✈︎ 30 running...")
	})
	tr.add(10, func() {
		log.Println("✈︎ 10 running...")
	})
	tr.add(100, func() {
		log.Println("✈︎ 100 running...")
	})

	time.Sleep(time.Hour)

}

type task struct {
	round    int
	callback taskFn
}

type (
	taskFn   func()
	taskList []*task
)

type timer struct {
	*time.Ticker
	sync.RWMutex
	sync.WaitGroup

	size      int
	cnt       int
	taskQueue map[int]taskList
}

func newTimer(size int) *timer {
	t := &timer{}
	t.taskQueue = make(map[int]taskList, size)
	t.size = size
	t.run()
	return t
}

func (t *timer) add(delay int, fn taskFn) {
	t.Lock()
	defer t.Unlock()

	round := delay / t.size
	index := delay % t.size
	log.Println("add task:", round, index)
	t.taskQueue[index] = append(t.taskQueue[index], &task{round: round, callback: fn})
}

func (t *timer) run() {
	t.Ticker = time.NewTicker(time.Second)
	go func() {
		for range t.Ticker.C {
			log.Println("time wheel is loop...", t.cnt)
			t.execute()
		}
	}()
}

func (t *timer) execute() {
	t.RLock()
	defer t.RUnlock()

	if res, ok := t.taskQueue[t.cnt]; ok {
		for k := 0; k < len(res); k++ {
			if res[k].round <= 0 {
				// 执行该函数
				res[k].callback()
				// 删除该函数
			} else {
				res[k].round--
			}
		}
	}

	t.cnt++
	if t.cnt == t.size {
		t.cnt = 0
	}

}

func (t *timer) stop() {
	if t.Ticker != nil {
		t.Ticker.Stop()
	}
}
