package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	bytes, err := base64.StdEncoding.DecodeString("eyJsdiI6MTExLCJwaGFzZSI6MjIyLCJzdGFyIjoyMjN9")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	var data = make(map[string]interface{})
	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(err)
	}

	fmt.Println("data map => ", data)

	/*
		output:
		{"lv":111,"phase":222,"star":223}
		data map =>  map[lv:111 phase:222 star:223]
	*/

}

type record struct {
	id     uint32
	parent []uint32
}

type records []record

func (rs records) search(id uint32) record {
	for _, r := range rs {
		if r.id == id {
			return r
		}
	}
	return record{}
}

type node struct {
	no   uint32
	prev []*node
	next []*node
}

type nodes []*node

type manager struct {
}

func source() []record {

	return nil
}

func checkDeadLockByDAG(records []record) (string, error) {

	return "", nil
}

func checkDeadLockByRecursion(records records) error {

	var stack = NewStack(4)

	var recursion func(rec record) error
	recursion = func(rec record) error {
		if stack.Contains(rec.id) {
			return fmt.Errorf("dead lock error:router:{%s}->entry point:{%d} ", stack, rec.id)
		}
		stack.Push(rec.id)
		for _, child := range rec.parent {
			if err := recursion(records.search(child)); err != nil {
				return err
			}
			stack.Pop()
		}
		return nil
	}

	for _, rec := range records {
		stack.Clear()
		if err := recursion(rec); err != nil {
			return err
		}
	}
	return nil
}

func checkDeadLockByUnionFindSet(records records) error {
	// 仅仅适用于单个分叉
	var (
		queue = NewQueue(4)
		stack = NewStack(4)
	)

	unionFindSet := func(begin uint32) error {
		queue.EnQueue(begin)
		stack.Push(begin)
		for queue.Len() > 0 {
			if value, ok := queue.DeQueue(); ok {
				if stack.Contains(value) {
					return fmt.Errorf("dead lock error:router:{%s}->entry point:{%d} ", stack, value)
				}
				stack.Push(value)
				for _, child := range records.search(value).parent {
					queue.EnQueue(child)
				}
			}
		}
		return nil
	}

	for _, v := range records {
		if err := unionFindSet(v.id); err != nil {
			return err
		}
		stack.Clear()
	}
	return nil
}

type Stack struct {
	values []uint32
	set    map[uint32]struct{}
}

func NewStack(size int) *Stack {
	if size == 0 {
		size = 8
	}
	return &Stack{
		values: make([]uint32, 0, size),
		set:    make(map[uint32]struct{}, size),
	}
}

func (s *Stack) Len() int {
	return len(s.values)
}

func (s *Stack) Push(value uint32) {
	s.values = append(s.values, value)
	s.set[value] = struct{}{}
}

func (s *Stack) Pop() (uint32, bool) {
	if len(s.values) == 0 {
		return 0, false
	}
	top := len(s.values) - 1
	value := s.values[top]
	s.values = s.values[:top]
	delete(s.set, value)
	return value, true
}

func (s Stack) Contains(value uint32) bool {
	_, ok := s.set[value]
	return ok
}

func (s *Stack) Clear() {
	if len(s.values) > 0 {
		s.values = s.values[:0]
	}
	s.set = make(map[uint32]struct{}, len(s.set)>>2)
}

func (s Stack) String() string {
	if len(s.values) > 0 {
		router := make([]string, len(s.values))
		for k, v := range s.values {
			router[k] = strconv.FormatUint(uint64(v), 10)
		}
		return strings.Join(router, "->")
	}
	return ""
}

type Queue struct {
	values []uint32
}

func NewQueue(size int) *Queue {
	if size == 0 {
		size = 4
	}
	return &Queue{
		values: make([]uint32, 0, size),
	}
}

func (q *Queue) Len() int {
	return len(q.values)
}

func (q *Queue) EnQueue(value uint32) {
	q.values = append(q.values, value)
}

func (q *Queue) DeQueue() (uint32, bool) {
	if len(q.values) > 0 {
		value := q.values[0]
		q.values = q.values[1:]
		return value, true
	}
	return 0, false
}
