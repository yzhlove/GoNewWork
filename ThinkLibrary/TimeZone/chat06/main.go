package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"think-library/TimeZone/chat06/assest"
	"time"
	_ "time/tzdata"
)

func main() {
	zones, err := getZones()
	if err != nil {
		panic(err)
	}

	now := time.Now()

	for _, zone := range zones {
		loc, err := time.LoadLocation(zone)
		if err != nil {
			panic(fmt.Sprintf("load zone [%s] error [%v]", zone, err))
		}

		fmt.Println(fmt.Sprintf("zone:%s \t\t\t %s", zone, now.In(loc).Format(time.RFC3339)))
	}
	fmt.Println("len ====> ", len(zones))
}

func getZones() ([]string, error) {

	var zones = make([]string, 0, 128)
	scan := bufio.NewScanner(bytes.NewBuffer(assest.ZoneData))

	var count = 3

	for scan.Scan() {
		text := scan.Text()
		if strings.HasPrefix(text, "#") {
			continue
		}
		for k, v := range strings.SplitN(text, "\t", count+1) {
			if k == count-1 {
				zones = append(zones, v)
				break
			}
		}
	}

	return zones, nil
}
