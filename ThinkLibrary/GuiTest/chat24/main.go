package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"time"
)

func main() {

	view, msgCh := Draw()
	tag := make(chan struct{})

	go func() {
		var index = 0
		for {
			index++
			var str = randomString(rand.Intn(200))
			select {
			case <-tag:
				msgCh <- Item{index: index, msg: str}
			}
			time.Sleep(time.Millisecond * 1000)
		}
	}()

	startBtn := widget.NewButton("START", func() {
		close(tag)
	})
	startBtn.Importance = widget.HighImportance

	MainWindow(func(win fyne.Window) fyne.CanvasObject {
		return container.NewBorder(nil, startBtn, nil, nil, view)
	})
}

func Draw() (*widget.List, chan Item) {

	var data = make([]Item, 0, 32)
	var msgCh = make(chan Item, 32)

	view := widget.NewList(func() int {
		return len(data)
	}, func() fyne.CanvasObject {
		icon := widget.NewIcon(theme.CancelIcon())
		richText := widget.NewTextGrid()
		richText.ShowLineNumbers = true
		return container.NewBorder(nil, nil, icon, nil, richText)
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		if widgets, ok := object.(*fyne.Container); ok {
			if len(widgets.Objects) >= 2 && len(data) > id {
				item := data[id]
				richText := widgets.Objects[0].(*widget.TextGrid)
				icon := widgets.Objects[1].(*widget.Icon)
				if item.status {
					icon.SetResource(theme.ConfirmIcon())
				} else {
					icon.SetResource(theme.CancelIcon())
				}
				richText.SetText(fmt.Sprintf("(%d)%s", item.index, item.msg))
			}
		}
	})

	go func() {
		for msg := range msgCh {
			data = append(data, msg)
			view.ScrollToBottom()
		}
	}()

	return view, msgCh
}

func MainWindow(fn func(mw fyne.Window) fyne.CanvasObject) {
	app := app.New()
	win := app.NewWindow("HelloTest")
	win.SetContent(fn(win))
	win.Resize(fyne.NewSize(1000, 600))
	win.ShowAndRun()
}

type Item struct {
	status bool
	index  int
	msg    string
}

var str = "ABCDEFGHIGKLMNOPQRSTUVWXYZ"

func randomString(size int) string {
	var strs []rune
	for i := 0; i < size; i++ {
		strs = append(strs, rune(str[rand.Intn(len(str))]))
	}
	return string(strs)
}
