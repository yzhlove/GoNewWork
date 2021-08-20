package main

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := encode("RedisTestChat.zip", "/Users/yostar/workSpace/GoNewWork/RedisTestChat");err != nil {
		panic(err)
	}
}

func encode(dst, src string) error {
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(f)
	if err = filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		w, err := zw.Create(obtained(src, path))
		if err != nil {
			return err
		}
		rf, err := os.Open(path)
		if err != nil {
			return err
		}
		_, err = io.Copy(w, rf)
		return err
	}); err != nil {
		return err
	}
	return zw.Close()
}

func obtained(src, path string) string {
	str := strings.Split(src, string(os.PathSeparator))
	prefix := strings.Join(str[:len(str)-1], string(os.PathSeparator)) + string(os.PathSeparator)
	return strings.Replace(path, prefix, "", 1)
}
