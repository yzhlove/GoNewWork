package main

import (
	"log"
	"math/rand"
)

type item struct {
	id uint32
	w  int32
}

var (
	items     = []item{{1, 10}, {2, 10}, {3, 20}, {4, 10}, {5, 35}, {6, 15}}
	itemsDept = []item{{1, 10}, {2, 10}, {4, 10}, {5, 35}, {6, 15}}
)

func main() {

	c := obtained()

	a := make(map[uint32]uint32, 6)
	b := make(map[uint32]uint32, 6)

	for i := 0; i < 100000; i++ {

		n := <-c

		at := take(n, 100)
		bt := take(n, 80)

		a[drop(at)]++
		b[dropDept(bt)]++
	}

	show(a)
	log.Println("------------------------------------")
	show(b)

}

func show(s map[uint32]uint32) {
	for k, v := range s {
		log.Printf("id:%d \tcount:%d  \trand:%0.f%% \n", k, v, (float32(v)/100000)*100)
	}
}

func obtained() chan float32 {

	c := make(chan float32, 10)
	go func() {
		for {
			c <- rand.Float32()
		}
	}()
	return c
}

func dropDept(number int32) uint32 {
	for _, a := range itemsDept {
		if a.w > number {
			return a.id
		}
		number -= a.w
	}
	return 0
}

func drop(number int32) uint32 {
	for _, a := range items {
		if a.id == 3 {
			number -= a.w
			continue
		}
		if a.w > number {
			return a.id
		}
		number -= a.w
	}
	return 0
}

func take(f float32, max int32) int32 {
	return int32(float32(max) * f)
}
