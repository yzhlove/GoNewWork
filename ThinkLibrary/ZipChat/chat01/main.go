package main

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
)

func main() {
	buf := compress([]byte("hello world++++++!"))

	// 创建一个文件用来装 zip的bytes.Buffer
	f, err := os.Create("local.zip")
	if err != nil {
		panic(err)
	}

	// 将zip的buf复制到文件
	io.Copy(f, buf)

	f.Sync()
	f.Close()
}

// data 是需要写入到zip的数据
func compress(data []byte) *bytes.Buffer {

	buf := bytes.NewBuffer([]byte{})

	zw := zip.NewWriter(buf)
	defer func() {
		zw.Flush()
		zw.Close()
	}()

	// 创建一个临时文件，文件名随意，主要是需要文件的fileInfo信息
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		// zip写入完成之后删除该文件
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

	// data 为实际写入的数据
	if _, err := io.Copy(w, bytes.NewReader(data)); err != nil {
		panic(err)
	}

	return buf
}
