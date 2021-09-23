package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func TopLayout(path string) fyne.CanvasObject {
	label := widget.NewLabel(" ")
	root := widget.NewIcon(theme.HomeIcon())
	pathLabel := widget.NewLabel(path)
	pathLabel.Alignment = fyne.TextAlignLeading
	pathLabel.TextStyle = fyne.TextStyle{Bold: true}
	return container.NewHBox(label, root, pathLabel)
}

func LeftLayout(event func(selectd string)) fyne.CanvasObject {
	return ListView(event)
}

func CenterCSVTool(event func()) fyne.CanvasObject {
	label := widget.NewLabel("zehgsfbjhsdbhsfbvjhbshvbs")
	btn := widget.NewButton("ok", event)
	btn.Importance = widget.MediumImportance
	return widget.NewCard("xlsx2csv tool", "", container.NewBorder(nil, btn, nil, nil, label))
}

func CenterUnzip(event func()) fyne.CanvasObject {

	return nil
}

func CenterLayout() fyne.CanvasObject {
	return CenterCSVTool(func() {})
}
