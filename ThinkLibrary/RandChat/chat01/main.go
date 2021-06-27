package main

import (
	"log"
	"math/rand"
)

type item struct {
	id uint32
	w  int32
}

func main() {

	drop()

}

func drop() {

	items := []item{{1, 10}, {2, 10}, {3, 20}, {4, 10}, {5, 35}, {6, 15}}

	fn := func() uint32 {
		ws := rand.Int31n(100)
		for _, a := range items {
			if a.w > ws {
				return a.id
			}
			ws -= a.w
		}
		return 0
	}

	ds := make(map[uint32]uint32, 6)

	for i := 0; i < 100000; i++ {
		ds[fn()]++
	}

	show(ds)

}

func notDrop() {

}

func show(s map[uint32]uint32) {
	for k, v := range s {
		log.Printf("id:%d \tcount:%d  \trand:%0.f%% \n", k, v, (float32(v)/100000)*100)
	}
}
