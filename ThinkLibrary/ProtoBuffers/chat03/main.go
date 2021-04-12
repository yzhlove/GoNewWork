package main

import "fmt"

//zigzag编码

func main() {

	var a int32 = -150
	e := encode(a)
	d := decode(e)
	fmt.Printf("encode:%v decode:%v \n", e, d)
}

func encode(n int32) int32 {
	return (n << 1) ^ (n >> 31)
}

func decode(n int32) int32 {
	return (n >> 1) ^ -(n & 1)
}
