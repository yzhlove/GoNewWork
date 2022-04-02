package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {

	test()

	//if err := assest.LoadData(); err != nil {
	//	panic(err)
	//}
	//
	//now := time.Now()
	//
	//for _, zone := range assest.Get() {
	//	loc, err := time.LoadLocationFromTZData(zone, assest.Match(zone))
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	strtime := now.In(loc).Format(time.RFC3339)
	//	fmt.Println("time => ", strtime)
	//}

}

func test() {

	var data = []byte(`
time =>  2022-04-02T04:41:47+02:00
time =>  2022-04-01T22:41:47-04:00
time =>  2022-04-02T03:41:47+01:00
time =>  2022-04-02T07:41:47+05:00
time =>  2022-04-02T02:41:47Z
time =>  2022-04-02T09:41:47+07:00
time =>  2022-04-02T07:41:47+05:00
time =>  2022-04-02T15:41:47+13:00
time =>  2022-04-02T11:41:47+09:00
time =>  2022-04-02T07:41:47+05:00
time =>  2022-04-02T03:41:47+01:00
`)

	scan := bufio.NewScanner(bytes.NewBuffer(data))
	for scan.Scan() {
		text := scan.Text()
		if len(text) == 0 {
			continue
		}

		if idx := strings.LastIndexAny(text, "+-Z"); idx != -1 {
			fmt.Println(fmt.Sprintf("text:%s \tsuffix:UTC%s", text, text[idx:]))
		}

	}

}
