package main

import (
	"github.com/Unknwon/goconfig"
)

func main() {

	root := "/Users/yostar/.toolbox_config.ini"

	c, err := goconfig.LoadConfigFile(root)
	if err != nil {
		panic(err)
	}

	goconfig.SaveConfigFile(c, root)

	//res := c.GetKeyList("ServerList")
	//fmt.Println("res => ", res)
	//
	//ret, err := c.GetSection("ServerList")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("ret --> ", ret)

}
