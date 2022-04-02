package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"think-library/TimeZone/chat09/assest"
	"time"
)

var zonesDef = []string{
	"Asia/Shanghai",                  // 中国上海
	"Asia/Urumqi",                    // 中国新疆
	"Asia/Tokyo",                     // 日本东京
	"Asia/Seoul",                     // 韩国首尔
	"America/Adak",                   // 美国阿达克
	"Pacific/Honolulu",               // 太平洋/檀香山
	"America/Anchorage",              // 美国/安克雷奇
	"America/Juneau",                 // 美国/朱诺
	"America/Metlakatla",             // 美国/梅特拉卡特拉
	"America/Nome",                   // 美国/诺姆
	"America/Sitka",                  // 美国/锡特卡
	"America/Yakutat",                // 美国/亚库塔特
	"America/Los_Angeles",            // 美国/洛杉矶
	"America/Boise",                  // 美国/博伊西
	"America/Denver",                 // 美国/丹佛
	"America/Phoenix",                // 美国/凤凰城
	"America/Chicago",                // 美国/芝加哥
	"America/Indiana/Knox",           // 美国/印第安纳/诺克斯
	"America/Indiana/Tell_City",      // 美国/印第安纳/泰尔城
	"America/Menominee",              // 美国/梅诺米尼
	"America/North_Dakota/Beulah",    // 美国/北达科他/比尤拉
	"America/North_Dakota/Center",    // 美国/北达科他/中心
	"America/North_Dakota/New_Salem", // 美国/北达科他/新塞勒姆
	"America/Detroit",                // 美国/底特律
	"America/Indiana/Indianapolis",   // 美国/印第安纳/印第安纳波利斯
	"America/Indiana/Marengo",        // 美国/印第安纳/马伦戈
	"America/Indiana/Petersburg",     // 美国/印第安纳/彼得斯堡
	"America/Indiana/Vevay",          // 美国/印第安纳/韦韦
	"America/Indiana/Vincennes",      // 美国/印第安纳/万塞讷
	"America/Indiana/Winamac",        // 美国/印第安纳/威纳马克
	"America/Kentucky/Louisville",    // 美国/肯塔基/路易斯维尔
	"America/Kentucky/Monticello",    // 美国/肯塔基/蒙蒂塞洛
	"America/New_York",               // 美国/纽约
}

func main() {

	//test()

	if err := assest.LoadData(); err != nil {
		panic(err)
	}
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

	now := time.Now()
	for _, zone := range zonesDef {
		loc, err := time.LoadLocationFromTZData(zone, assest.Match(zone))
		if err != nil {
			panic(fmt.Sprintf("zone %s error: %s", zone, err))
		}

		fmt.Println("zone ", zone, "\t\t  time ", now.In(loc).Format(time.RFC3339))
	}

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
