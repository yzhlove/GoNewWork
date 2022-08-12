package main

import (
	"log"
	"unsafe"
)

func main() {

	val := int(0)
	p := &val

	type K struct{}

	a := K{}

	log.Println(unsafe.Sizeof(val))
	log.Println(unsafe.Sizeof(p))
	log.Println(unsafe.Sizeof(a))
	log.Printf("point address: %p", &a)

	var str = "i.mooc"
	log.Printf("string %d ", unsafe.Sizeof(str))
	log.Printf("string %d ", unsafe.Sizeof(&str))

}
