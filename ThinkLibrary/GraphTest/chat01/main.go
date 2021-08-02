package main

import (
	"fmt"
)

//数据结构:图

func main() {
	DFS()
	fmt.Println()
	callDFS()
}

var vector = []string{"A", "B", "C", "D", "E"}
var matrix = [5][5]int{
	{0, 1, 1, 1, 0},
	{0, 0, 1, 0, 1},
	{0, 0, 0, 0, 1},
	{0, 0, 1, 0, 0},
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
				break¬
			}
		}
		if i == len(vector) {
			stack = stack[:len(stack)-1]
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
