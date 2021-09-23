package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	res := app.New()
	w := res.NewWindow("TestDemo")
	btnDemo(w)
	w.Resize(fyne.NewSize(1000, 800))
	w.ShowAndRun()
}

func btnDemo(w fyne.Window) {
	label := widget.NewLabel("label str")
	btn1 := widget.NewButton("btn1", func() {
		label.SetText("btn1 onclick")
	})

	btn1.Importance = widget.HighImportance

	btn2 := widget.NewButton("btn2", func() {
		label.SetText("btn2 onclick")
	})

	btn2.Importance = widget.MediumImportance
	btn2.IconPlacement = widget.ButtonIconLeadingText

	btn3 := widget.NewButton("btn3", func() {
		label.SetText("btn3 onclick")
	})

	btn3.IconPlacement = widget.ButtonIconTrailingText
	btn3.Importance = widget.LowImportance

	c := container.NewVBox(label, btn1, btn2, btn3)
	w.SetContent(c)
}
