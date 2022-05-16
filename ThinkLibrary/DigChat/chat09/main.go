package main

import (
	"fmt"
)

func main() {

	nodes, recMap := create(data())
	nodes = nodes

	show(recMap, recMap[5])
}

type node struct {
	id   uint32
	prev []node
	next []node
}

func create(records []record) ([]node, map[uint32]record) {
	nodes := make([]node, 0, len(records))
	recMap := make(map[uint32]record, len(records))
	for _, rec := range records {
		nodes = append(nodes, node{id: rec.id})
		recMap[rec.id] = rec
	}
	return nodes, recMap
}

func show(collections map[uint32]record, rec record) {

	var loop func(rec record)
	var stack []string
	use := make(map[string]struct{})

	find := func(id string) bool {
		for _, s := range stack {
			if s == id {
				return true
			}
		}
		return false
	}

	loop = func(rec record) {
		// 访问过的点不在访问
		if _, ok := use[fmt.Sprintf("%d", rec.id)]; ok {
			fmt.Printf("value = %d %t [end] \n", rec.id, len(rec.prev) == 0)
			return
		} else {
			fmt.Printf("value = %d %t \n", rec.id, len(rec.prev) == 0)
		}

		// head node
		if len(rec.prev) == 0 {
			if !find(fmt.Sprintf("%d", rec.id)) {
				fmt.Println("------------> push ", rec.id)
				stack = append(stack, fmt.Sprintf("%d", rec.id))
				use[fmt.Sprintf("%d", rec.id)] = struct{}{}
			}
			return
		}

		for _, pid := range rec.prev {
			if ret, ok := collections[pid]; ok {
				loop(ret)
			}
		}

		for len(stack) > 0 {
			top := stack[len(stack)-1]
			fmt.Printf("======> %d insert: %s \n", rec.id, top)
			stack = stack[:len(stack)-1]
		}

		if !find(fmt.Sprintf("%d", rec.id)) {
			fmt.Println("------------> push ", rec.id)
			stack = append(stack, fmt.Sprintf("%d", rec.id))
			use[fmt.Sprintf("%d", rec.id)] = struct{}{}
		}
	}

	loop(rec)
}

type record struct {
	id   uint32
	prev []uint32
}

func data() []record {
	return []record{
		{10, []uint32{8, 6, 7, 9}},
		{9, []uint32{8, 7, 6}},
		{8, []uint32{6, 7}},
		{7, []uint32{6, 5}},
		{6, []uint32{4, 3, 5}},
		{5, []uint32{4, 33}},
		{4, []uint32{3, 2, 1}},
		{3, []uint32{1, 2}},
		{2, []uint32{1}},
		{1, []uint32{}},
		{33, []uint32{32}},
		{32, []uint32{31}},
		{31, []uint32{30}},
		{30, []uint32{1}},
	}
}
