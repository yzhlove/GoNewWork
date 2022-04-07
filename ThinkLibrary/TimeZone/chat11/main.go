package main

import (
	"fmt"
	"time"
)

func main() {
	test()
}

func test() {

	//sahnghai := "Asia/Shanghai"
	newYork := "America/New_York"

	newYorkLoc, err := time.LoadLocation(newYork)
	if err != nil {
		panic(err)
	}

	tm, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-03-13 02:25:25", newYorkLoc)
	if err != nil {
		panic(err)
	}

	fmt.Println("2022-03-13 02:25:25  ===> ", tm.Format("2006-01-02 15:04:05"))
	// 2022-03-13 02:25:25  ===>  2022-03-13 01:25:25
}

func test2() {
	shangLoc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	value := "2022-03-13T14:25:25+08:00"
	shangTime, err := time.ParseInLocation(time.RFC3339, value, shangLoc)
	if err != nil {
		panic(err)
	}

	newYorkLock, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	fmt.Println(shangTime.In(newYorkLock).Format("2006-01-02 15:04:05"))
	// "2022-03-13T15:25:25+08:00" ===> 2022-03-13 03:25:25
	// "2022-03-13T14:25:25+08:00" ===> 2022-03-13 01:25:25
}
