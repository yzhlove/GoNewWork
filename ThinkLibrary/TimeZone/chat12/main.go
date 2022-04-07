package main

import (
	"fmt"
	"time"
)

const (
	newYork  = "America/New_York"
	shanghai = "Asia/Shanghai"

	format = "2006-01-02 15:04:05"
)

func main() {
	test4()

}

func test1() {
	loc, err := time.LoadLocation(newYork)
	if err != nil {
		panic(err)
	}

	timestr := "2022-03-13 03:15:00"
	now, err := time.ParseInLocation(format, timestr, loc)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("isDST %t src %s dst %s ", now.IsDST(), timestr, now.Format(format)))
}

func test2() {

	shanghaiLoc, err := time.LoadLocation(shanghai)
	if err != nil {
		panic(err)
	}

	newYorkLoc, err := time.LoadLocation(newYork)
	if err != nil {
		panic(err)
	}

	timestr := "2022-11-06 02:10:00"
	now, err := time.ParseInLocation(format, timestr, shanghaiLoc)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("shanghai %s %s", timestr, now.Format(format)))

	tnow, err := time.ParseInLocation(format, timestr, newYorkLoc)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("newYork %t %s %s", tnow.IsDST(), timestr, tnow.Format(format)))

}

func test3() {

	newYorkLoc, err := time.LoadLocation(newYork)
	if err != nil {
		panic(err)
	}

	shanghaiLoc, err := time.LoadLocation(shanghai)
	if err != nil {
		panic(err)
	}

	timestr := "2022-11-06 13:00:00"
	now, err := time.ParseInLocation(format, timestr, shanghaiLoc)
	if err != nil {
		panic(err)
	}

	addNow := now.Add(time.Hour)

	fmt.Printf("A %t %t %s %s\n",
		addNow.IsDST(),
		addNow.In(newYorkLoc).IsDST(),
		addNow.Format(format),
		addNow.In(newYorkLoc).Format(format))

	fmt.Printf("B %t %t %s %s\n",
		now.IsDST(),
		now.In(newYorkLoc).IsDST(),
		now.Format(format),
		now.In(newYorkLoc).Format(format))

	subNow := now.Add(-time.Hour)

	fmt.Printf("C %t %t %s %s\n",
		subNow.IsDST(),
		subNow.In(newYorkLoc).IsDST(),
		subNow.Format(format),
		subNow.In(newYorkLoc).Format(format))

}

func test4() {

	newYorkLoc, err := time.LoadLocation(newYork)
	if err != nil {
		panic(err)
	}

	shanghaiLoc, err := time.LoadLocation(shanghai)
	if err != nil {
		panic(err)
	}

	timestr := "2022-11-06 01:00:00"
	now, err := time.ParseInLocation(format, timestr, newYorkLoc)
	if err != nil {
		panic(err)
	}

	fmt.Println(timestr, ",", now.Format(format))

	fmt.Println(fmt.Sprintf("%t %s %s",
		now.IsDST(), timestr, now.In(shanghaiLoc).Format(format)))

}
