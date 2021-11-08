package main

import (
	"bytes"
	"fmt"
	"go.uber.org/dig"
	"os"
)

func main() {

	//type Student struct {
	//	Name string
	//	Age  int
	//}
	//
	//type Rep struct {
	//	dig.Out
	//	StudentList []*Student `group:"stu,flatten"` //flatten的意思是，底层把组表示成[]*Student，如果不加flatten会表示成[][]*Student 完整示例
	//}
	//
	//NewUser := func(name string, age int) func() Rep {
	//	return func() Rep {
	//		r := Rep{}
	//		r.StudentList = append(r.StudentList, &Student{name, age})
	//		return r
	//	}
	//}

	type Student struct {
		Name string
		Age  int
	}
	type Rep struct {
		dig.Out
		StudentList []*Student `group:"stu,flatten"`
	}
	NewUser := func(name string, age int) func() Rep {
		return func() Rep {
			r := Rep{}
			r.StudentList = append(r.StudentList, &Student{
				Name: name,
				Age:  age,
			})
			return r
		}
	}

	container := dig.New()
	if err := container.Provide(NewUser("tom", 3)); err != nil {
		panic(err)
	}
	if err := container.Provide(NewUser("jerry", 1)); err != nil {
		panic(err)
	}

	//c := dig.New()
	//
	//if err := c.Provide(NewUser("tom", 3), dig.Group("stu")); err != nil {
	//	panic(err)
	//}
	//if err := c.Provide(NewUser("jerry", 1), dig.Group("stu")); err != nil {
	//	panic(err)
	//}

	type inParams struct {
		dig.In
		StudentList []*Student `group:"stu"`
	}

	err := container.Invoke(func(params inParams) {
		for _, u := range params.StudentList {
			fmt.Println(u.Name, "---", u.Age)
		}
	})

	if err != nil {
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
