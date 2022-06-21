package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"sync"
	"time"
)

func main() {

	listView, msgCh := GetListView()
	tag := make(chan struct{})
	once := sync.Once{}

	go func() {
		var index = 0
		for {
			index++
			var str = randomString(rand.Intn(100))
			select {
			case <-tag:
				msgCh <- Item{index: index, msg: str}
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()

	startBtn := widget.NewButton("START", func() {
		once.Do(func() {
			close(tag)
		})
	})
	startBtn.Importance = widget.HighImportance

	MainWindow(func(win fyne.Window) fyne.CanvasObject {
		return container.NewBorder(nil, startBtn, nil, nil, listView)
	})

}

func MainWindow(callback func(win fyne.Window) fyne.CanvasObject) {

	application := app.New()
	window := application.NewWindow("Test")

	window.SetContent(callback(window))
	window.Resize(fyne.NewSize(1000, 600))
	window.ShowAndRun()
}

func GetListView() (*widget.List, chan Item) {

	var data = make([]Item, 0, 32)
	var msgCh = make(chan Item, 32)

	listView := widget.NewList(func() int {
		return len(data)
	}, func() fyne.CanvasObject {
		tagIcon := widget.NewLabel("")
		tagIcon.TextStyle = fyne.TextStyle{Italic: true, Bold: true}
		msgLabel := widget.NewMultiLineEntry()
		msgLabel.TextStyle = fyne.TextStyle{Bold: true}
		msgLabel.Wrapping = fyne.TextWrapBreak
		msgLabel.Disable()
		return container.NewBorder(nil, nil, tagIcon, nil, msgLabel)
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		if widgets, ok := object.(*fyne.Container); ok {
			if len(widgets.Objects) >= 2 && len(data) > id {
				item := data[id]
				widgets.Objects[0].(*widget.Entry).SetText(item.msg)
				widgets.Objects[1].(*widget.Label).SetText(fmt.Sprintf("%d", item.index))
			}
		}
	})

	go func() {
		for msg := range msgCh {
			data = append(data, msg)
			listView.ScrollToBottom()
		}
	}()

	return listView, msgCh
}

type Item struct {
	index int
	msg   string
}

var str = "ABCDEFGHIGKLMNOPQRSTUVWXYZ"

func randomString(size int) string {
	var strs []rune
	for i := 0; i < size; i++ {
		strs = append(strs, rune(str[rand.Intn(len(str))]))
	}
	return string(strs)
}
