package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"time"
)

func main() {

	application := app.New()
	window := application.NewWindow("Test")

	d1 := dialog.NewInformation("A", "what are you doing?", window)
	d1.Show()

	go func() {
		d2 := dialog.NewInformation("B", "it`s Ok", window)
		d2.Show()
		time.Sleep(time.Second * 5)
		d2.Hide()
	}()

	window.Resize(fyne.NewSize(800, 350))
	window.ShowAndRun()

}
