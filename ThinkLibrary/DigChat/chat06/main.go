package main

import (
	"fmt"
	"go.uber.org/dig"
)

func main() {
	type Student struct {
		dig.Out
		Name string
		Age  *int `option:"true"` //允许一个字段为nil
	}

	c := dig.New()
	_ = c.Provide(func() Student {
		return Student{Name: "tom"}
	})

	_ = c.Invoke(func(name string, age *int) {
		if age == nil {
			fmt.Println("age is nil")
		} else {
			fmt.Printf("age: %d ", age)
		}
	})

}
