package main

import (
	"fmt"
	"time"
)

func main() {

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	tokyoLoc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	fmt.Println("current time:", time.Now().In(loc).Format(time.RFC3339))
	fmt.Println("current time:", time.Now().In(tokyoLoc).Format(time.RFC3339))

}
