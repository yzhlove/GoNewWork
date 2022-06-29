package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math/rand"
)

func main() {

	view, msgCh := Draw()

	btn := widget.NewButton("START", func() {
		go generate(msgCh, 10)
	})
	btn.Importance = widget.HighImportance

	MainWindow(func(win fyne.Window) fyne.CanvasObject {
		return container.NewBorder(nil, btn, nil, nil, view)
	})

}

var runes = []rune(`ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`)
var runelen = len(runes)

func character(size int) string {

	var strs = make([]rune, 0, size)
	for i := 0; i < size; i++ {
		strs = append(strs, runes[rand.Intn(runelen)])
	}
	return string(strs)
}

func generate(msgCh chan string, count int) {
	for i := 0; i < count; i++ {
		msgCh <- character(rand.Intn(100))
	}
}

func MainWindow(callback func(win fyne.Window) fyne.CanvasObject) {

	application := app.New()
	window := application.NewWindow("Test")

	window.SetContent(callback(window))
	window.Resize(fyne.NewSize(600, 400))
	window.ShowAndRun()
}

func Draw() (*widget.List, chan string) {

	var data = make([]string, 0, 1024)
	var msgCh = make(chan string, 512)

	view := widget.NewList(func() int {
		return len(data)
	}, func() fyne.CanvasObject {

		richText := widget.NewRichText()
		richText.Wrapping = fyne.TextWrapBreak
		size := richText.MinSize()
		richText.Resize(fyne.NewSize(size.Width, size.Height*2))

		return richText
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		richText := object.(*widget.RichText)
		richText.Segments = richText.Segments[:0]
		richText.Segments = append(richText.Segments, &widget.TextSegment{Text: data[id]})
		richText.Refresh()
	})

	go func() {
		for msg := range msgCh {
			data = append(data, msg)
			view.ScrollToBottom()
		}
	}()
	return view, msgCh
}
