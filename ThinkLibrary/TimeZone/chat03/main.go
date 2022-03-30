package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	f, err := os.Open("zone.tab")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	buf := strings.Builder{}
	var count = 3
	var zone string

	now := time.Now()

	for scan.Scan() {
		if strings.HasPrefix(scan.Text(), "#") {
			continue
		}
		text := scan.Text()
		strs := strings.SplitN(text, "\t", count+1)
		if len(strs) < count {
			panic(fmt.Sprintf("%s error", text))
		}

		buf.Reset()
		for k, tv := range strs {
			switch k {
			case 0:
				buf.WriteString(" code:" + tv)
			case 1:
				buf.WriteString(" location:" + tv)
			case 2:
				buf.WriteString(" zone:" + tv)
				zone = tv
			case 3:
				buf.WriteString(" desc:" + tv)
			default:
				panic("invalid str")
			}
		}

		loc, err := time.LoadLocation(zone)
		if err != nil {
			panic(fmt.Sprintf(" load zone [%s] error: %v", zone, err))
		}

		fmt.Println(fmt.Sprintf("[%s] \t\t [%s]", zone, now.In(loc).Format(time.RFC3339)))
	}

}
