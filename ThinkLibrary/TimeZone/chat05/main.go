package main

import (
	_ "embed"
	"fmt"
)

//go嵌入静态资源

//go:embed zone.tab
var s string

func main() {
	fmt.Println(s)
}
