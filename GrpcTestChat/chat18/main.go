package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("MaxUint32 => ", math.MaxUint32)
	fmt.Println("MaxUint32 - 9 = ", math.MaxUint32-9)
	fmt.Println("int32(MaxUint32 - 9) = ", tr(math.MaxUint32-9))

	/*
		MaxUint32 =>  4294967295
		MaxUint32 - 9 =  4294967286
		int32(MaxUint32 - 9) =  -10
	*/
}

func tr(bb uint32) int32 {
	return int32(bb)
}
