package main

import (
	"context"
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
	tr.add(0, func() {
		log.Println("✈︎ 0 running...")
	})

	time.Sleep(time.Hour)

}

type taskFunc func()

type task struct {
	round    int
	callback taskFunc
}

type taskManager struct {
	taskQueue []*task
}

func (t *taskManager) add(tk *task) {
	t.taskQueue = append(t.taskQueue, tk)
}

func (t *taskManager) remove(index int) {
	if index >= 0 && len(t.taskQueue) > index {
		t.taskQueue = append(t.taskQueue[:index], t.taskQueue[index+1:]...)
	}
}

func (t *taskManager) iterate(callback func(k int, tk *task) error) error {
	for i := 0; i < len(t.taskQueue); i++ {
		if err := callback(i, t.taskQueue[i]); err != nil {
			return err
		}
	}
	return nil
}

type taskExecute struct {
	sync.RWMutex
	context.Context

	taskCh chan taskFunc
}

func newTaskExecute(ctx context.Context, size int) *taskExecute {
	if ctx == nil {
		ctx, _ = context.WithCancel(context.Background())
	}
	execute := &taskExecute{Context: ctx, taskCh: make(chan taskFunc, size)}
	execute.start()
	return execute
}

func (t *taskExecute) add(fn taskFunc) {
	t.taskCh <- fn
}

func (t *taskExecute) start() {
	go func() {
		for {
			select {
			case <-t.Done():
				return
			case callback := <-t.taskCh:
				go callback()
			}
		}
	}()
}

type timer struct {
	context.Context
	context.CancelFunc
	*time.Ticker
	sync.Mutex
	*taskExecute
	size    int
	cursor  int
	manager map[int]*taskManager
}

func newTimer(size int) *timer {
	t := &timer{size: size}
	t.Context, t.CancelFunc = context.WithCancel(context.Background())
	t.taskExecute = newTaskExecute(t.Context, 128)
	t.manager = make(map[int]*taskManager, size)
	t.run()

	return t
}

func (t *timer) next() {
	if t.cursor == t.size {
		t.cursor = 0
	}
	t.cursor++
}

func (t *timer) add(delay int, fn taskFunc) {
	t.Lock()
	defer t.Unlock()

	round := delay / t.size
	index := delay % t.size

	if t.manager[index] == nil {
		t.manager[index] = &taskManager{}
	}

	t.manager[index].add(&task{round: round, callback: fn})
}

func (t *timer) run() {
	t.Ticker = time.NewTicker(time.Second)
	go func() {
		for range t.Ticker.C {
			log.Println("timer cursor: ", t.cursor)
			t.execute()
		}
	}()
}

func (t *timer) execute() {
	t.Lock()
	defer t.Unlock()

	if manager, ok := t.manager[t.cursor]; ok {
		manager.iterate(func(k int, tk *task) error {
			log.Println("iterator task:", tk.round)
			if tk.round <= 0 {
				t.taskExecute.add(tk.callback)
				manager.remove(k)
			} else {
				tk.round--
			}
			return nil
		})
	}

	t.next()
}
