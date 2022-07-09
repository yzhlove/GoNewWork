package main

import (
	"context"
	"sync"
	"time"
)

func main() {

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
	return &taskExecute{Context: ctx, taskCh: make(chan taskFunc, size)}
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
	t := &timer{}
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
			t.execute()
		}
	}()
}

func (t *timer) execute() {
	t.Lock()
	defer t.Unlock()

	if manager, ok := t.manager[t.cursor]; ok {
		manager.iterate(func(k int, tk *task) error {
			if tk.round == 0 {
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
