package main

import (
	"fmt"
	"think-library/TimeZone/chat07/assest"
	"time"
)

func main() {

	zones := assest.GetZones()
	data, err := assest.LoadZoneData()
	if err != nil {
		panic(err)
	}

	//assest.ReadZoneDataZip()

	now := time.Now()
	for _, zone := range zones {
		if r, ok := data[zone]; ok {

			loc, err := time.LoadLocationFromTZData(zone, r)
			if err != nil {
				panic(err)
			}

			fmt.Println("zone", zone, "\tcurrent time ", now.In(loc).Format(time.RFC3339))

		} else {
			fmt.Println("------------------------------------------------------------- not found zone ===> ", zone)
		}
	}

}
