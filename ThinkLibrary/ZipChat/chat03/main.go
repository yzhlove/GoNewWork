package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	code    = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	codeLen = len(code)
)

func main() {

	source := rand.NewSource(12345678)
	rd := rand.New(source)

	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}

	defer func() {
		file.Sync()
		file.Close()
	}()

	var count = 10000000
	start := time.Now()
	for i := 0; i < count; i++ {
		userId := rd.Uint64()
		str := getChar(rd, 11)
		file.WriteString(fmt.Sprintf("%s,%d,1\r\n", str, userId))
	}

	log.Printf("witer over, time:%dms \n ", time.Now().Sub(start).Milliseconds())
}

func getChar(rd *rand.Rand, number int) string {
	str := make([]rune, number)
	for k, v := range rd.Perm(codeLen)[:number] {
		str[k] = rune(code[v])
	}
	return string(str)
}
