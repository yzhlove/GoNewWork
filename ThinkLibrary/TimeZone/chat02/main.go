package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	bytes, err := os.ReadFile("tzdata2022a/zone.tab")
	if err != nil {
		panic(err)
	}

	loc, err := time.LoadLocationFromTZData("Asia/Shanghai", bytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Now().In(loc).Format(time.RFC3339))

}
