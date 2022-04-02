package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"strings"
)

//go:embed zone.txt
var a []byte

//go:embed zone_transform.txt
var b []byte

func main() {

	ra := read(a)
	rb := read(b)
	if err := write(ra, rb); err != nil {
		panic(err)
	}
	fmt.Println("Ok")
}

func read(data []byte) (r []string) {
	scan := bufio.NewScanner(bytes.NewBuffer(data))
	for scan.Scan() {
		r = append(r, scan.Text())
	}
	return
}

func write(aa, bb []string) error {
	if len(aa) == len(bb) {
		var sb strings.Builder
		for k := range aa {
			sb.WriteString(fmt.Sprintf("%s\t%s\n", aa[k], bb[k]))
		}

		f, err := os.Create("zone_merge.txt")
		if err != nil {
			return err
		}
		defer f.Close()

		f.WriteString(sb.String())
		f.Sync()

		return nil
	}
	return errors.New("inconsistent")
}
