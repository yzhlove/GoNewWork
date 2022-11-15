package main

import (
	"fmt"
	"strings"
)

type Inputer interface {
	Set(any)
	Get() any
}

type MiddlewareFunc func(input Inputer) Inputer

type Middleware struct {
	handles []MiddlewareFunc
}

func NewMiddle(handles ...MiddlewareFunc) *Middleware {
	return &Middleware{
		handles: handles,
	}
}

func (m *Middleware) Run(input Inputer) Inputer {
	for _, handfunc := range m.handles {
		input = handfunc(input)
	}
	return input
}

type Values struct {
	params []any
}

func (v *Values) Set(x any) {
	v.params = append(v.params, x)
}

func (v *Values) Get() any {
	return v.params
}

func (v *Values) String() string {
	var sb strings.Builder
	for _, vt := range v.params {
		sb.WriteString(fmt.Sprintf("%T->%+v\n", vt, vt))
	}
	return sb.String()
}

func SetMap(data map[string]string) MiddlewareFunc {
	return func(input Inputer) Inputer {
		for k, v := range data {
			input.Set(fmt.Sprintf("%s:%v", k, v))
		}
		return input
	}
}

func SetSlice(strs []string) MiddlewareFunc {
	return func(input Inputer) Inputer {
		input.Set(strings.Join(strs, "."))
		return input
	}
}

func main() {
	middle := NewMiddle(
		SetMap(map[string]string{"x": "123", "y": "456"}),
		SetSlice([]string{"hello", "world"}),
	)

	x := &Values{}
	ret := middle.Run(x)
	fmt.Println(ret)
}
