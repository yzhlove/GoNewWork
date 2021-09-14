package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

//压缩和解压缩

func main() {
	if err := zipEncode("RedisTestChat.zip", "./RedisTestChat"); err != nil {
		panic(err)
	}
}

func zipEncode(dst, src string) error {

	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fw.Close()

	zw := zip.NewWriter(fw)
	defer func() {
		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		zf, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		//zf.Name = strings.TrimPrefix(path, string(filepath.Separator))
		//if info.IsDir() {
		//	zf.Name += "/"
		//}

		w, err := zw.CreateHeader(zf)
		if err != nil {
			return err
		}

		if !zf.Mode().IsRegular() {
			return nil
		}

		fr, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fr.Close()

		_, err = io.Copy(w, fr)
		if err != nil {
			return err
		}

		fmt.Printf("成功压缩文件:%s", path)
		return nil
	})
}

func zipDecode(dst, src string) error {

	zr, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer zr.Close()

	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}

	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)

		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			continue
		}

		fr, err := file.Open()
		if err != nil {
			return err
		}

		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
		if err != nil {
			return err
		}

		_, err = io.Copy(fw, fr)
		if err != nil {
			return err
		}
		fmt.Printf("成功解压 %s ", path)
		fw.Close()
		fr.Close()
	}
	return nil
}
