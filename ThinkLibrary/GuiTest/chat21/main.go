package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strings"
)

func main() {

	str := strings.Repeat("A", 200)

	a := widget.NewLabel(str)
	a.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	//a.Wrapping = fyne.TextWrapOff

	b := widget.NewLabel(str)
	b.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	//b.Wrapping = fyne.TextTruncate

	c := widget.NewLabel(str)
	c.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	c.Wrapping = fyne.TextWrapBreak

	d := widget.NewLabel(str)
	d.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	//d.Wrapping = fyne.TextWrapWord

	MainWindow(func(win fyne.Window) fyne.CanvasObject {
		return container.NewBorder(c, nil, nil, nil)
	})

}

func MainWindow(callback func(win fyne.Window) fyne.CanvasObject) {

	application := app.New()
	window := application.NewWindow("Test")

	window.SetContent(callback(window))
	window.Resize(fyne.NewSize(600, 400))
	window.ShowAndRun()
}
