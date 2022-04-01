package assest

import (
	"bufio"
	"bytes"
	_ "embed"
	"strings"
)

const (
	maxSplit = 4
	zoneLoc  = 2
)

//go:embed zone.tab
var zonetab []byte

func GetZones() []string {
	scan := bufio.NewScanner(bytes.NewBuffer(zonetab))
	zones := make([]string, 0, 64)
	for scan.Scan() {
		text := scan.Text()
		if strings.HasPrefix(text, "#") {
			continue
		}
		for k, v := range strings.SplitN(text, "\t", maxSplit) {
			if k == zoneLoc {
				zones = append(zones, v)
				break
			}
		}
	}
	return zones
}
