package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"strings"
)

func main() {

	c := app.NewWithID("yostar")
	c.Settings().SetTheme(&MyTheme{})
	w := c.NewWindow("服务端工具集")

	topContent := TopLayout("/Users/yurisa/Develop/GoWork/src/WorkSpace/GoNewWork/ThinkLibrary/GuiTest/chat11/main.go")

	xlsxBox := XlsxBox(func(label *widget.Label) {
		label.SetText("打表成功")
	})

	updateBox := UpdateZipBox(func(label *widget.Label, s []string) {
		label.SetText(strings.Join(s, "\n"))
	})

	gachaBox := GachaBox(func(label *widget.Label) {
		f, err := ioutil.ReadFile("/Users/yostar/workSpace/GoNewWork/ThinkLibrary/GuiTest/chat11/layoutComponent.go")
		if err != nil {
			label.SetText(err.Error())
		} else {
			label.SetText(string(f))
		}
	})

	contentBoxs := []fyne.CanvasObject{xlsxBox, updateBox, gachaBox}
	for i := range ListItems {
		ListItems[i].Widget = contentBoxs[i]
	}

	mainShow := container.NewGridWithColumns(1, contentBoxs...)

	ctrl := LeftLayout(func(id widget.ListItemID) {
		fmt.Println("id -> ", id, " name ->", ListItems[id].Name)
		for i := range ListItems {
			if ListItems[i].Widget != nil {
				if i == id {
					ListItems[i].Widget.Show()
				} else {
					ListItems[i].Widget.Hide()
				}
			}
		}
	})

	cnet := container.NewHSplit(ctrl, mainShow)
	cnet.SetOffset(0.2)

	copyright := BottomLayout()

	w.SetContent(container.NewBorder(topContent, copyright, nil, nil, cnet))
	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()
}
