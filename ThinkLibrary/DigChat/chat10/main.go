package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	nodes, records := create(data())
	if err := build(nodes, records); err != nil {
		panic(err)
	}

	//for _, node := range nodes {
	//	fmt.Println("==============================")
	//	fmt.Println("=== node.id ", node.id)
	//	for _, prev := range node.prev {
	//		fmt.Println("== node.prev ", prev.id)
	//	}
	//	for _, next := range node.next {
	//		fmt.Println("== node.next ", next.id)
	//	}
	//	fmt.Println("==============================")
	//}

	//showPrev(nodes[910005])
	fmt.Println()
	showNext(nodes[1])
}

func showPrev(start *node) {
	fmt.Print(start.id, " -> ")
	for _, prev := range start.prev {
		showPrev(prev)
	}
}

func showNext(start *node) {
	fmt.Print(start.id, " -> ")
	for _, next := range start.next {
		showNext(next)
	}
}

func create(data []record) (map[uint32]*node, map[uint32]record) {
	nodes := make(map[uint32]*node)
	records := make(map[uint32]record)

	for _, x := range data {
		n := &node{id: x.id}
		if len(x.prev) == 0 {
			n.isTop = true
		}
		nodes[n.id] = n
		records[x.id] = x
	}

	return nodes, records
}

func build(nodes map[uint32]*node, records map[uint32]record) error {

	var iterator func(rec record) error
	var stack []uint32
	var visit = make(map[uint32]struct{})
	var router = make([]string, 0, 8)

	checkDeadLock := func(strs []string) bool {
		c := make(map[string]struct{}, len(strs))
		for _, x := range strs {
			if _, ok := c[x]; ok {
				return true
			}
			c[x] = struct{}{}
		}
		return false
	}

	check := func(x uint32) bool {
		for _, t := range stack {
			if x == t {
				return true
			}
		}
		return false
	}

	iterator = func(x record) error {
		fmt.Println("--> x.id", x.id)
		//入栈
		router = append(router, strconv.FormatUint(uint64(x.id), 10))
		if checkDeadLock(router) {
			return fmt.Errorf("dead lock error:entry point: %d , chanins: %s", x.id, strings.Join(router, "->"))
		}

		// 已经被访问过的点不在访问
		if _, ok := visit[x.id]; ok {
			return nil
		}

		for _, id := range x.prev {
			if rt, ok := records[id]; ok {
				if err := iterator(rt); err != nil {
					return err
				}
				// 出栈
				if len(router) > 0 {
					router = router[:len(router)-1]
				}
			}
		}

		node, ok := nodes[x.id]
		if !ok {
			return fmt.Errorf("current ndoe [%d] missing", x.id)
		}
		if len(stack) > 0 {
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				if tnode, ok := nodes[top]; ok {
					//fmt.Println(x.id, "->", tnode.id)
					fmt.Println(tnode.id, "->", x.id)
					fmt.Println()
					node.setPrev(tnode)
					tnode.setNext(node)
				} else {
					return fmt.Errorf("node [%d] missing", top)
				}
				stack = stack[:len(stack)-1]
			}
		}

		if !check(x.id) {
			stack = append(stack, x.id)
			visit[x.id] = struct{}{}
		}

		return nil
	}

	for _, ret := range records {
		stack = stack[:0]
		router = router[:0]
		visit = make(map[uint32]struct{}, len(records))
		if err := iterator(ret); err != nil {
			return err
		}
		fmt.Println("--------------------------------------")
	}

	// set Next
	//for _, node := range nodes {
	//	for _, c := range node.prev {
	//		c.setNext(node)
	//	}
	//}

	return nil
}

type node struct {
	id    uint32
	prev  []*node
	next  []*node
	isTop bool
}

func (n *node) setNext(nt *node) {
	for _, k := range n.next {
		if k.id == nt.id {
			return
		}
	}
	n.next = append(n.next, nt)
}

func (n *node) setPrev(nt *node) {
	for _, k := range n.prev {
		if k.id == nt.id {
			return
		}
	}
	n.prev = append(n.prev, nt)
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

/*
1 2 3 4 30
2 3 4
3 4 6
4 5 6
5 6
32 33
31 32
30 31
*/

/*
1 30 2 3 4
2 3 4
3 4 6
4 5 6
5 6 7
6 7 8 9 10
7 8 9 10
8 9
8 10
32 33
31 32
30 31



*/

/*
func data() []record {
	return []record{
		{5, []uint32{4}},
		{4, []uint32{3, 2}},
		{3, []uint32{2}},
		{2, []uint32{1, 5}},
		{1, []uint32{}},
	}
}
*/

/*
func data() []record {
	return []record{
		{910001, []uint32{}},
		{910002, []uint32{910001}},
		{910003, []uint32{910002}},
		{910004, []uint32{910003, 910002}},
		{910005, []uint32{910003, 910001}},
	}
}
*/
