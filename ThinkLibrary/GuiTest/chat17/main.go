package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	testNotify()

}

func testNotify() {
	MainWindow(func(mw fyne.Window) fyne.CanvasObject {

		return widget.NewButton("test", func() {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "test notify",
				Content: "this is send notify...",
			})
		})
	})
}

func MainWindow(fn func(mw fyne.Window) fyne.CanvasObject) {
	app := app.New()
	win := app.NewWindow("HelloTest")
	win.SetContent(fn(win))
	win.Resize(fyne.NewSize(1000, 600))
	win.ShowAndRun()
}
