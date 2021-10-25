package main

import (
	"fmt"
	"log"
	"reflect"
	"unsafe"
)

func main() {

	var a int
	var b uint16
	var c uint32
	var d uint64

	var s = []interface{}{a, b, c, d}

	log.Println("a size:", unsafe.Sizeof(a))
	log.Println("a size:", unsafe.Sizeof(b))
	log.Println("a size:", unsafe.Sizeof(c))
	log.Println("a size:", unsafe.Sizeof(d))
	log.Println("a size:", unsafe.Sizeof(s))

	log.Println()

	for _, k := range s {
		log.Println(reflect.TypeOf(k).Size())
		fmt.Println("-> ", unsafe.Sizeof(k))
	}

}
