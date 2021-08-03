package main

import (
	"fmt"
)

//数据结构:图

func main() {
	//DFS()
	//fmt.Println()
	//callDFS()
	//ddfs("A")
	checkDFS("A")
}

var vector = []string{"A", "B", "C", "D", "E"}
var matrix = [5][5]int{
	{0, 1, 1, 1, 0},
	{0, 0, 1, 0, 1},
	{0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0},
}

/*
DAG图
	A	B	C	D	E
A	0	1	1	1	0
B	0	0	1	0	1
C	0	0	0	0	1
D	0	0	1	0	0
E	0	0	0	1	0
*/

func DFS() {
	heads := make([]string, 0, 1)
	//剪枝，找出入读为0的点
Next:
	for j := 0; j < len(vector); j++ {
		for i := 0; i < len(vector); i++ {
			if matrix[i][j] == 1 {
				break Next
			}
		}
		heads = append(heads, vector[j])
	}

	fmt.Println("heads ===> ", heads)

	for _, v := range heads {
		dfs(v)
	}
}

func dfs(str string) {
	x := find(str)
	stack := make([]int, 0, len(vector))
	rept := make(map[int]int, len(vector))

	//初始化一个节点
	stack = append(stack, x)
	rept[x]++
	var i int
	fmt.Printf(" (%s,%d) ", str, x)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		for i = 0; i < len(vector); i++ {
			if matrix[top][i] == 1 && rept[i] == 0 {
				fmt.Printf(" (%s,%d) ", vector[i], i)
				stack = append(stack, i)
				rept[i]++
				break
			}
		}
		if i == len(vector) {
			stack = stack[:len(stack)-1]
		}
	}

}

func ddfs(str string) {
	x := find(str)
	stack := make([]int, 0, len(vector))
	rept := make(map[int]int, len(vector))

	stack = append(stack, x)
	rept[x]++
	fmt.Printf(" (%s,%d) ", vector[x], x)
	var i int

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		for i = 0; i < len(vector); i++ {
			//是否可达
			if matrix[top][i] == 1 {
				fmt.Printf(" (%s,%d) ", vector[i], i)
				rept[i]++
				//入环点
				if rept[i] > 1 {
					panic(fmt.Sprintf("dead lock by (%s,%d) ,count: %d ", vector[i], i, rept[i]))
				}
				stack = append(stack, i)
				break
			}
		}
		if i == len(vector) {
			stack = stack[:len(stack)-1]
		}
	}
}

func checkDFS(str string) {
	type node struct {
		name    string //节点名
		inEdge  int    //入度
		outEdge int    //出度
	}

	nodes := make([]node, len(vector))
	for i := 0; i < len(vector); i++ {
		var in, out int
		for j := 0; j < len(vector); j++ {
			if matrix[j][i] == 1 {
				in++
			}
			if matrix[i][j] == 1 {
				out++
			}
		}
		nodes[i] = node{name: vector[i], inEdge: in, outEdge: out}
	}

	fmt.Printf("nodes ===> %+v \n", nodes)

	search := func(str string) int {
		for k, n := range nodes {
			if n.name == str {
				return k
			}
		}
		return -1
	}

	x := search(str)
	stack := make([]int, 0, len(vector))
	edge := make(map[int]int, len(vector))
	visit := make(map[int]struct{}, len(vector))

	stack = append(stack, x)
	visit[x] = struct{}{}
	fmt.Printf(" [%s] ", vector[x])

	var i int
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		fmt.Printf("(↓ %s->%d)", vector[top], edge[top])
		for i = edge[top]; i < len(vector); i++ {
			if matrix[top][i] == 1 {
				if _, ok := visit[i]; !ok {
					edge[top] = i + 1
					fmt.Printf(" [%s] ", vector[i])
					stack = append(stack, i)
					visit[i] = struct{}{}
					break
				} else {
					panic(fmt.Sprintf("[deal lock by (%s,%d)]", vector[i], i))
				}
			}
		}
		if i == len(vector) {
			t := stack[len(stack)-1]
			//fmt.Printf("(↑ %s)", vector[t])
			stack = stack[:len(stack)-1]
			delete(visit, t)
		}
	}

}

func find(str string) int {
	for k, v := range vector {
		if v == str {
			return k
		}
	}
	return -1
}

func callDFS() {
	idx := find("A")
	rept := make(map[int]int, len(vector))
	dfs2(idx, rept)
}

func dfs2(point int, rept map[int]int) {
	fmt.Printf(" (%s,%d) ", vector[point], point)
	rept[point]++
	for i := 0; i < len(vector); i++ {
		if rept[i] == 0 && matrix[point][i] == 1 {
			dfs2(i, rept)
		}
	}
}
