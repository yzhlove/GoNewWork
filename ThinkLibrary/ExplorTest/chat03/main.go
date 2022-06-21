package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	//s := `/home/yurisa/abc/a/b/c/d`
	s := "C:\\Documents\\Newsletters\\Summer2018.pdf"
	//// 下面这句用于 Windows 系统
	//s = filepath.FromSlash(s)
	//fmt.Println(s) // /a/b/c/d 或 \a\b\c\d
	// 下面这句用于 Windows 系统
	s = filepath.ToSlash(s)
	fmt.Println(s) // /a/b/c/d
}
