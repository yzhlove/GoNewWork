package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/dtylman/gowd"
)

//go:embed main.html
var s string

//go:embed output.log
var str []byte

func main() {
	body, err := gowd.ParseElement("<h1>Hello world</h1>", nil)
	if err != nil {
		panic(err)
	}
	p := body.AddElement(gowd.NewElement("p"))
	btn := p.AddElement(gowd.NewElement("button"))
	btn.SetText("Click me")
	btn.OnEvent(gowd.OnClick, btnClicked)
	gowd.Run(body)
}

func btnClicked(sender *gowd.Element, event *gowd.EventElement) {
	sender.SetText("Clicked!")
}

func main1() {

	buffer := bytes.NewBuffer(str)

	fmt.Println("s -> ",s)

	body, err := gowd.ParseElement(s, nil)
	if err != nil {
		buffer.WriteString(err.Error())
		panic(err)
	}

	btn := body.Find("btn_a")
	if btn != nil {
		btn.SetText("Click Me")
		btn.OnEvent(gowd.OnClick, btnClicked)
	}

	if err := gowd.Run(body); err != nil {
		buffer.WriteString(err.Error())
		panic(err)
	}
}
