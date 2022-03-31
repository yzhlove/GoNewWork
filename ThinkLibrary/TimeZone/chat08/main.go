package main

import (
	"fmt"
	"think-library/TimeZone/chat08/assest"
	"time"
)

func main() {

	if err := assest.LoadData(); err != nil {
		panic(err)
	}

	now := time.Now()

	for _, v := range assest.Get() {
		data := assest.Match(v)
		if data == nil {
			panic(fmt.Sprintf("not found zone %v data", v))
		}

		loc, err := time.LoadLocationFromTZData(v, data)
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("zone:%s \t==> \t time:%s", v, now.In(loc).Format(time.RFC3339)))
	}

}
