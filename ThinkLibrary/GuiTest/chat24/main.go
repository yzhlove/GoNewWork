package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {

	MainWindow(func(win fyne.Window) fyne.CanvasObject {

		cust := dialog.NewCustom("test", "dsbfhasbvhadfbhv", widget.NewLabel("this is label"), win)
		btn := widget.NewButton("open", func() {
			cust.Show()
		})

		return container.NewBorder(nil, btn, nil, nil)
	})

}

func MainWindow(callback func(win fyne.Window) fyne.CanvasObject) {

	application := app.New()
	window := application.NewWindow("Test")

	window.SetContent(callback(window))
	window.Resize(fyne.NewSize(600, 400))
	window.ShowAndRun()
}
