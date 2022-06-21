package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path1 := "a/b/c.txt"
	path2 := "/a/b/c.txt"

	fmt.Printf("join:%s , slash:%s \n", filepath.Join(dir, path1), filepath.ToSlash(filepath.Join(dir, path1)))
	fmt.Printf("join:%s , slash:%s \n", filepath.Join(dir, path2), filepath.ToSlash(filepath.Join(dir, path2)))
}

/*
join:/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/ThinkLibrary/a/b/c.txt , slash:/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/ThinkLibrary/a/b/c.txt
join:/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/ThinkLibrary/a/b/c.txt , slash:/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/ThinkLibrary/a/b/c.txt

*/
