package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

/*
type RichTextSegment interface {
    Inline() bool
    Textual() string
    Update(fyne.CanvasObject)
    Visual() fyne.CanvasObject
    Select(pos1 fyne.Position, pos2 fyne.Position)
    SelectedText() string
    Unselect()
}
*/

func main() {
	//testProcess1()
	testProcess2()
}

func MainWindow(fn func(mw fyne.Window) fyne.CanvasObject) {
	app := app.New()
	win := app.NewWindow("HelloTest")
	win.SetContent(fn(win))
	win.Resize(fyne.NewSize(1000, 600))
	win.ShowAndRun()
}

func testProcess1() {
	MainWindow(func(mw fyne.Window) fyne.CanvasObject {

		process := widget.NewProgressBarInfinite()
		process.Resize(fyne.NewSize(1000, 10))

		rich := widget.NewRichText()
		scroll := container.NewVScroll(rich)

		go func() {

			for i := 0; i < 100; i++ {
				rich.Segments = append(rich.Segments, &widget.TextSegment{Text: fmt.Sprintf("\nhello world [%d]", i+1), Style: widget.RichTextStyleInline})
				rich.Refresh()
				scroll.ScrollToBottom()
				scroll.Refresh()
				time.Sleep(time.Second)
			}
		}()

		return container.NewBorder(process, nil, nil, nil, scroll)
	})
}

func testProcess2() {

	var data = make([]string, 0, 100)
	var index = 0
	for ; index < 10; index++ {
		data = append(data, fmt.Sprintf("hello %d", index+1))
	}

	MainWindow(func(mw fyne.Window) fyne.CanvasObject {

		list := widget.NewList(func() int {
			return len(data)
		}, func() fyne.CanvasObject {
			return widget.NewLabel("")
		}, func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(data[id])
		})

		go func() {
			for ; index < 200; index++ {
				data = append(data, fmt.Sprintf("hello %d", index+1))
				list.ScrollToBottom()
				list.Refresh()
				time.Sleep(time.Millisecond * 200)
				fmt.Println("index value ---> ", index)
			}
		}()

		return list
	})
}
