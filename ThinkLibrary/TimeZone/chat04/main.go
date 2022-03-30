package main

import (
	"fmt"
	"os"
)

func main() {

	if err := os.Setenv("ZONEINFO", "tzdata.zip"); err != nil {
		panic(fmt.Sprintf("设置环境变量错误."))
	}

}
