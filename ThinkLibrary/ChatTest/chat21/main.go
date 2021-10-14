package main

import "fmt"

func main() {

	m := Map{}
	a := Array{}
	for i := 0; i < 2; i++ {
		m.With(fmt.Sprintf("object:%d", i), fmt.Sprintf("value:%d", i+1))
		a.With(fmt.Sprintf("array-%d", i))
	}
	m.Show()
	fmt.Println("----------")
	m.Clear()
	m.Show()
	fmt.Println("----------")
	a.Show()
}

type Map map[string]string

func (m Map) With(key, value string) {
	m[key] = value
}

func (m *Map) Clear() {
	t := *m
	t = make(map[string]string)
	*m = t
}

func (m Map) Show() {
	for k, v := range m {
		fmt.Println(fmt.Sprintf("k:%v v:%v ", k, v))
	}
}

type Array []string

func (a *Array) With(value string) {
	t := *a
	t = append(t, value)
	*a = t
}

func (a Array) Show() {
	for _, v := range a {
		fmt.Println("v -> ", v)
	}
}
