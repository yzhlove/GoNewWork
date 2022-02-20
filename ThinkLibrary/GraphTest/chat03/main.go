package main

import (
	"fmt"
	"sort"
)

func main() {

	p := presents{}
	p.load()

	p.show()
}

/*
Id		Num
1001	1
1002	2
1003	3
1004	4
1005	5
*/

type record struct {
	Id  uint32
	Num int
}

type info struct {
	record
	disassemble int
}

type ins []info

func (i ins) Len() int {
	return len(i)
}

func (i ins) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i ins) Less(x, y int) bool {
	return i[x].disassemble <= i[y].disassemble
}

func (t info) show() string {
	return fmt.Sprintf("Id<%d>,Num<%d>,disassemble<%d>", t.Id, t.Num, t.disassemble)
}

func create() []record {
	return []record{
		//{1000, 0}, // 0级 默认
		{1000, 1}, // 1级
		{1000, 2}, // 2级
		{1000, 3}, // 3级
		{1000, 4}, // 4级
		{1000, 5}, // 5级
	}
}

type presents struct {
	data ins
}

func (p *presents) load() {
	p.data = make([]info, 0, 4)
	records := create()
	for k := range records {
		p.data = append(p.data, info{records[k], 0})
	}
	p.disassemble()

	return
}

func (p *presents) disassemble() {
	for k, v := range p.data {
		if k == 0 {
			p.data[k].disassemble = 2 // 0星 + 1星
		} else {
			p.data[k].disassemble = p.data[k-1].disassemble + v.Num
		}
	}
	sort.Sort(p.data)
}

func (p presents) show() {
	for _, v := range p.data {
		fmt.Println(v.show())
	}
}

type material struct {
	star   int // 星际
	number int // 数量
}

func (p presents) Update(star int, ms []material) (int, material, error) {

	if star >= 0 && star <= len(p.data) {
		// 如果已经是最高等级，则不需要升级
		if star == len(p.data) {
			return 0, material{}, nil
		}
		if len(ms) > 0 {
			// 当前星级提供的卡片
			var ownQty int
			if star > 0 {
				ownQty = p.data[star-1].disassemble
			}
			// 所有的素材能分解出来的卡片
			for _, v := range ms {
				ownQty += p.data[v.star-1].disassemble * v.number
			}

			index := sort.Search(len(p.data), func(i int) bool {
				return p.data[i].disassemble >= ownQty
			})
			fmt.Println("ownQty => ", ownQty, " index => ", index)
			// 如果没有找到
			if index == len(p.data) {
				return 0, material{}, fmt.Errorf("[presents] not found need upgrade")
			}

			return index + 1, material{0, p.data[index].disassemble - ownQty}, nil
		}
		return 0, material{}, nil
	}

	return 0, material{}, fmt.Errorf("[presents] star [%d] out of range", star)
}
