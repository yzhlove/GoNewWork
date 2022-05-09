package main

import (
	"fmt"
	"strings"
)

func main() {

	/*
		100 0
		101 100
		102 100
		103 101
		104 101
		105 101
		108 105
	*/

	var array [][2]int
	array = data()

	// 建立关系
	relation := make(map[int][]int)

	for _, v := range array {
		cnt, parent := v[0], v[1]
		relation[parent] = append(relation[parent], cnt)
	}

	nodes := make([]*node, 0, len(relation))
	head := &node{}
	nodes = append(nodes, head)

	nodes = build(head, nodes, relation)

	for _, k := range nodes {
		xn = 0
		stat(k)
		k.n = xn
	}

	for _, k := range nodes {
		fmt.Println("node k => ", k.id, k.n)
	}

	fmt.Println("------------------------------")
	show(head, 0)
}

type node struct {
	id   int
	next []*node
	n    int
}

func create(nodes []*node, id int) (*node, bool) {
	for _, v := range nodes {
		if v.id == id {
			return v, true
		}
	}
	return &node{id: id}, false
}

func build(n *node, nodes []*node, relation map[int][]int) []*node {
	xs, ok := relation[n.id]
	if !ok {
		return nodes
	}

	for _, x := range xs {
		tnode, exist := create(nodes, x)
		if !exist {
			nodes = append(nodes, tnode)
		}
		n.next = append(n.next, tnode)
		nodes = build(tnode, nodes, relation)
	}

	return nodes
}

var xn = 0

func stat(n *node) {
	xn++
	if len(n.next) == 0 {
		return
	}
	for _, tn := range n.next {
		stat(tn)
	}
}

var list []string

func show(n *node, dept int) {

	list = append(list, fmt.Sprintf("%d[%d]", n.id, n.n))

	for _, tn := range n.next {
		show(tn, dept+1)
	}

	if len(n.next) == 0 {
		if len(list) > 0 {
			fmt.Println(strings.Join(list, "->"))
		}
	}

	if len(list) > 0 {
		list = list[:len(list)-1]
	}
}

func data() [][2]int {
	return [][2]int{
		{100, 0},
		{101, 100},
		{102, 100},
		{103, 101},
		{104, 101},
		{105, 101},
		{106, 103},
		{107, 103},
		{108, 107},
	}
}
