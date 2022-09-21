package main

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
)

// golang zip 压缩

func main() {
	buf := compress([]byte("hello world++++++!"))
	f, err := os.Create("local.zip")
	if err != nil {
		panic(err)
	}
	io.Copy(f, buf)
	f.Sync()
	f.Close()
}

func compress(data []byte) *bytes.Buffer {

	buf := bytes.NewBuffer([]byte{})

	zw := zip.NewWriter(buf)
	defer func() {
		zw.Flush()
		zw.Close()
	}()

	// 可以创建临时文件 os.CreateTemp()
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		os.Remove(file.Name())
	}()

	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fh, err := zip.FileInfoHeader(fi)
	if err != nil {
		panic(err)
	}
	// 启用压缩
	fh.Method = zip.Deflate

	w, err := zw.CreateHeader(fh)
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(w, bytes.NewReader(data)); err != nil {
		panic(err)
	}

	return buf
}
