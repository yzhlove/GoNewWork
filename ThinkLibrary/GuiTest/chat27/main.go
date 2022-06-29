package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"strings"
)

func main() {

	rich := widget.NewRichText(&widget.TextSegment{Text: "AAA"})
	rich.Wrapping = fyne.TextWrapBreak
	//log.Printf("rich text size -> %+v ", rich.MinSize())

	var index = 0

	btn := widget.NewButton("A", func() {
		if index%2 == 0 {
			rich.Segments = rich.Segments[:0]
			rich.Segments = append(rich.Segments, &widget.TextSegment{Text: strings.Repeat("ABC", 30)})
			log.Printf("a  rich text size -> %+v ", rich.MinSize())
			rich.Refresh()
		} else {
			rich.Segments = rich.Segments[:0]
			rich.Segments = append(rich.Segments, &widget.TextSegment{Text: strings.Repeat("ABC", 2)})
			log.Printf("b  rich text size -> %+v ", rich.MinSize())
			rich.Refresh()
		}
		index++
	})

	MainWindow(func(win fyne.Window) fyne.CanvasObject {
		return container.NewBorder(nil, btn, nil, nil, rich)
	})

}

func MainWindow(callback func(win fyne.Window) fyne.CanvasObject) {

	application := app.New()
	window := application.NewWindow("Test")

	window.SetContent(callback(window))
	window.Resize(fyne.NewSize(400, 200))
	window.ShowAndRun()
}
