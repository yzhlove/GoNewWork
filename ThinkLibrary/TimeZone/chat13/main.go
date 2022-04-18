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

	str := "2022-03-13 02:10:00"

	york, err := time.LoadLocation(newYork)
	if err != nil {
		panic(err)
	}

	t, err := time.ParseInLocation(format, str, york)
	if err != nil {
		panic(err)
	}

	fmt.Println("==> ", t.Format(format), str)

	shang, err := time.LoadLocation(shanghai)
	if err != nil {
		panic(err)
	}

	fmt.Println("==> ", t.In(shang).Format(format), str)

	tt, err := time.ParseInLocation(format, str, shang)
	if err != nil {
		panic(err)
	}

	fmt.Println("==> ", tt.Format(format), str)
	fmt.Println("==> ", tt.In(york).Format(format), str)

}
