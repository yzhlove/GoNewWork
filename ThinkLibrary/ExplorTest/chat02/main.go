package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {

	data := []byte("hello world")

	buf := bytes.NewReader(data)
	str, err := fileByHash(buf)
	fmt.Printf("hash:%s %v", str, err)
	buf.Seek(0, 0)
	copyToStd(buf)

}

func fileByHash(reader io.Reader) (string, error) {
	hash := md5.New()
	_, err := io.Copy(hash, reader)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func copyToStd(reader io.Reader) error {
	_, err := io.Copy(os.Stdout, reader)
	return err
}
