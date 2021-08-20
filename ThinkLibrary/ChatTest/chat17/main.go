package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := encode("thinkLibrary.zip", "/Users/yostar/workSpace/GoNewWork/ThinkLibrary"); err != nil {
		panic(err)
	}
}

func encode(dst, src string) error {
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(f)
	buf := new(bytes.Buffer)

	if err = filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		name := obtained(src, path)
		w, err := zw.Create(name)
		if err != nil {
			return err
		}
		rf, err := os.Open(path)
		if err != nil {
			return err
		}
		hash := checkCode(rf)
		buf.WriteString(fmt.Sprintf("file:%s\tcheckCode:%s\n", name, hash))
		_, err = io.Copy(w, rf)
		return err
	}); err != nil {
		return err
	}
	verify := src + string(os.PathSeparator) + "verify.txt"
	w, err := zw.Create(obtained(src, verify))
	if err != nil {
		return err
	}
	if _, err = io.Copy(w, buf); err != nil {
		return err
	}
	return zw.Close()
}

func obtained(src, path string) string {
	str := strings.Split(src, string(os.PathSeparator))
	prefix := strings.Join(str[:len(str)-1], string(os.PathSeparator)) + string(os.PathSeparator)
	return strings.Replace(path, prefix, "", 1)
}

func checkCode(f io.ReadWriteCloser) string {
	h256 := sha256.New()
	if _, err := io.Copy(h256, f); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h256.Sum(nil))
}
