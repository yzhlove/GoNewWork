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

	//for _, v := range assest.Get() {
	//	data := assest.Match(v)
	//	if data == nil {
	//		panic(fmt.Sprintf("not found zone %v data", v))
	//	}
	//
	//	loc, err := time.LoadLocationFromTZData(v, data)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Println(fmt.Sprintf("zone:%s \t==> \t time:%s", v, now.In(loc).Format(time.RFC3339)))
	//}

	//fmt.Println(now.Format(time.RFC3339))

	shanghaiLoc, _ := time.LoadLocationFromTZData("Asia/Shanghai", assest.Match("Asia/Shanghai"))
	//fmt.Println(now.In(shanghaiLoc).Format(time.RFC3339))
	//fmt.Println(now.In(time.UTC).Format(time.RFC3339))

	tokyoLoc, _ := time.LoadLocationFromTZData("Atlantic/St_Helena", assest.Match("Atlantic/St_Helena"))
	fmt.Println(now.In(tokyoLoc).Format(time.RFC3339))

	tm, err := time.ParseInLocation("2006-01-02 15:03:04", "1991-05-20 00:10:00", shanghaiLoc)
	if err != nil {
		panic(err)
	}
	fmt.Println("1991-05-20 00:10:00 format ==> ", tm.Format(time.RFC3339))

}
