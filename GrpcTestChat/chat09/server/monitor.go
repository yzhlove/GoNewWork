package main

import "fmt"

type monitorType struct{}

type Monitor struct {
	data []string
}

func New() *Monitor {
	return &Monitor{data: make([]string, 0)}
}

func (m *Monitor) Push(strs ...string) {
	m.data = append(m.data, strs...)
}

func (m *Monitor) Clear() {
	m.data = m.data[:0]
}

func (m *Monitor) Show() {
	fmt.Println("monitor show:", m.data)
}
