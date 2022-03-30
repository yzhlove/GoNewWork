package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	bytes, err := os.ReadFile("/usr/share/zoneinfo/Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	loc, err := time.LoadLocationFromTZData("Asia/Tokyo", bytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Now().In(loc).Format(time.RFC3339))

}
