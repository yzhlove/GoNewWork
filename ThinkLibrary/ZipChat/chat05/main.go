package main

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"math"
	"strconv"
)

func main() {
	test2()

	var t uint16 = 0
	tb := make([]byte, 2)
	binary.BigEndian.PutUint16(tb, t)
	fmt.Println("crc32 = ", crc32.ChecksumIEEE(tb))

}

func test1() {
	var value uint64 = math.MaxUint64
	var b uint64 = 15
	var sum uint64
	for i := 0; i < 16; i++ {
		vv := value & b
		sum += vv
		value >>= 4
		fmt.Printf("i:%d  %v :%b \n", i, vv, vv)
	}
	fmt.Printf("sum =  %d \n", sum)
}

func test2() {

	var value uint64 = math.MaxUint64
	fmt.Printf("maxUint 64 = %d \n", value)
	var sum int
	str := strconv.FormatUint(value, 10)
	for _, v := range str {
		tv, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		sum += tv
	}
	fmt.Println("sum = ", sum)
	sum = 0

	for value > 0 {
		a := value % 10
		fmt.Printf("%d", a)
		sum += int(a)
		value /= 10
	}
	fmt.Println()
	fmt.Println("sum = ", sum)

}
