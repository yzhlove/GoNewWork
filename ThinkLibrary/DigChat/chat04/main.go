package main

import (
	"bytes"
	"fmt"
	"go.uber.org/dig"
	"os"
)

func main() {

	type Student struct {
		Name string
		Age  int
	}

	NewUser := func(name string, age int) func() *Student {
		return func() *Student {
			return &Student{Name: name, Age: age}
		}
	}

	container := dig.New()
	_ = container.Provide(NewUser("tom", 3), dig.Group("stu"))
	_ = container.Provide(NewUser("jerry", 1), dig.Group("stu"))

	type inParams struct {
		dig.In
		StudentList []*Student `group:"stu"` //group  和 name 不能同时使用
	}

	if err := container.Invoke(func(params inParams) {
		for _, stu := range params.StudentList {
			fmt.Println("stu => ", stu)
		}
	}); err != nil {
		panic(err)
	}

	showDot(container)

}

func showDot(c *dig.Container) {
	buf := bytes.Buffer{}
	if err := dig.Visualize(c, &buf); err != nil {
		panic(err)
	}
	f, err := os.Create("dig.dot")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(buf.Bytes())
}
