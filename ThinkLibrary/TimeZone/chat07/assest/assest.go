package assest

import (
	"archive/zip"
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"path"
	"strings"
)

const (
	maxSplit = 4
	zoneLoc  = 2
)

//go:embed zone.tab
var zoneData []byte

//go:embed tzdata.zip
var tzBytes []byte

func GetZones() []string {
	scan := bufio.NewScanner(bytes.NewBuffer(zoneData))
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

func ReadZoneDataZip() error {

	r := bytes.NewReader(tzBytes)

	reader, err := zip.NewReader(r, r.Size())
	if err != nil {
		return err
	}

	for _, f := range reader.File {
		fmt.Println(f.Name + "\t\t ===> " + path.Dir(f.Name))
	}

	fmt.Println("len ---> ", len(reader.File))

	return nil
}

func LoadZoneData() (map[string][]byte, error) {
	tr := bytes.NewReader(tzBytes)
	reader, err := zip.NewReader(tr, tr.Size())
	if err != nil {
		return nil, err
	}

	data := make(map[string][]byte, 128)
	for _, f := range reader.File {
		// 如果是文件夹
		if f.FileInfo().IsDir() {
			continue
		}
		// 打开文件
		stream, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("[%s] open error:[%v]", f.Name, err)
		}

		zdata, err := io.ReadAll(stream)
		if err != nil {
			return nil, fmt.Errorf("[%s] read error:[%v]", f.Name, err)
		}

		data[f.Name] = zdata
	}

	return data, err
}
