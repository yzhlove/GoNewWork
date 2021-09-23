package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {

	c := app.NewWithID("yostar")
	w := c.NewWindow("ServerToolBox")

	topContent := TopLayout("/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/ThinkLibrary/GuiTest/chat11/main.go")
	leftContent := LeftLayout(func(selectd string) {
		fmt.Println("str -> ", selectd)
	})

	centerContent := CenterLayout()
	cnet := container.NewHSplit(leftContent, centerContent)
	cnet.SetOffset(0.2)

	w.SetContent(container.NewBorder(topContent, nil, nil, nil, cnet))
	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()
}
