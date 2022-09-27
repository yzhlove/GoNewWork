package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"io"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()

	application := app.New()
	window := application.NewWindow("Test")

	labelText := widget.NewLabel("")
	labelText.Wrapping = fyne.TextWrapBreak

	//log.Println(string(data))
	btn := widget.NewButton("Button", func() {
		//newDialog(window, string(data))
		labelText.SetText(string(data))
	})

	window.SetContent(container.NewBorder(nil, btn, nil, nil, container.NewScroll(labelText)))
	window.Resize(fyne.NewSize(800, 350))
	window.ShowAndRun()
}

func newDialog(w fyne.Window, s string) {
	label := widget.NewLabel(s)
	label.Wrapping = fyne.TextWrapBreak

	wrap := container.NewGridWrap(fyne.NewSize(600, 150), container.NewScroll(label))

	dg := dialog.NewCustom("Msg:", "Close", wrap, w)
	dg.Show()
}
