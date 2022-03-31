package assest

import (
	"archive/zip"
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	tabName  = "zone.tab"
	maxSplit = 4
	locIndex = 2
)

var (
	errLoad   = errors.New("load tzdata.zip error")
	errNotTab = errors.New("load zone.tab file error")
)

//go:embed tzdata.zip
var tzdata []byte

type zoneData struct {
	zones   []string          // 所有的区
	records map[string][]byte // 所有区的数据
}

var data zoneData

func LoadData() error {
	r := bytes.NewReader(tzdata)
	reader, err := zip.NewReader(r, r.Size())
	if err != nil {
		return fmt.Errorf("load tzdata zip file error:%v", err)
	}

	data.records = make(map[string][]byte, 128)
	for _, f := range reader.File {
		if f.Mode().IsDir() {
			continue
		}
		rf, err := f.Open()
		if err != nil {
			return fmt.Errorf("open %s file error:%v", f.Name, err)
		}
		record, err := io.ReadAll(rf)
		if err != nil {
			return fmt.Errorf("read %s file error:%v", f.Name, err)
		}
		data.records[f.Name] = record
	}

	return data.generateZone()
}

func (z *zoneData) generateZone() error {
	if len(z.records) == 0 {
		return errLoad
	}

	if r, ok := z.records[tabName]; ok {
		z.zones = make([]string, 0, 128)
		scan := bufio.NewScanner(bytes.NewBuffer(r))

		for scan.Scan() {
			text := scan.Text()
			if strings.HasPrefix(text, "#") {
				continue
			}
			for k, v := range strings.SplitN(text, "\t", maxSplit) {
				if k == locIndex {
					z.zones = append(z.zones, v)
					break
				}
			}
		}

		z.cleanInvalidZone()
		return nil
	}

	return errNotTab
}

func (z *zoneData) cleanInvalidZone() {
	if len(z.zones) == 0 {
		return
	}

	match := func(zone string) bool {
		for _, v := range z.zones {
			if v == zone {
				return true
			}
		}
		return false
	}

	for k := range z.records {
		if match(k) {
			continue
		}
		delete(z.records, k)
	}

	return
}

func Get() []string {
	return data.zones
}

func Match(zone string) []byte {
	if len(data.records) > 0 {
		return data.records[zone]
	}
	return nil
}
