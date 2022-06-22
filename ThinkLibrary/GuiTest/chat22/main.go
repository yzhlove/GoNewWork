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

	canvs, msgCh := Draw()
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
		return container.NewBorder(nil, startBtn, nil, nil, canvs)
	})

}

func MainWindow(callback func(win fyne.Window) fyne.CanvasObject) {

	application := app.New()
	window := application.NewWindow("Test")

	window.SetContent(callback(window))
	window.Resize(fyne.NewSize(1000, 600))
	window.ShowAndRun()
}

func Draw() (fyne.CanvasObject, chan Item) {

	var msgCh = make(chan Item, 32)

	richText := widget.NewRichText()
	richText.Wrapping = fyne.TextWrapWord

	scroll := container.NewVScroll(richText)
	go func() {
		for msg := range msgCh {
			richText.Segments = append(richText.Segments,
				&widget.TextSegment{
					Text:  fmt.Sprintf("(%d)%s", msg.index, msg.msg),
					Style: widget.RichTextStyleCodeBlock,
				})
			richText.Refresh()
			scroll.ScrollToBottom()
		}
	}()

	return scroll, msgCh
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
