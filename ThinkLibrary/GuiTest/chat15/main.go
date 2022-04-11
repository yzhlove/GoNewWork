package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	c := app.NewWithID("12138")
	w := c.NewWindow("toolbox")

	aLabel := widget.NewLabel("please input time string")
	aEntry := widget.NewEntry()

	topLayout := container.NewBorder(
		nil, nil, aLabel, nil, aEntry)

	leftLabel := widget.NewLabel("left")
	leftSelect := widget.NewSelect([]string{"a", "b", "c"}, func(s string) {})
	rightLabel := widget.NewLabel("right")
	rightSelect := widget.NewSelect([]string{"a", "b", "c"}, func(s string) {})

	leftLayout := container.NewGridWithColumns(2,
		container.NewPadded(leftLabel), container.NewPadded(rightLabel))
	rightLayout := container.NewGridWithColumns(2,
		container.NewPadded(leftSelect), container.NewPadded(rightSelect))

	tipsLabel := widget.NewLabel("this is show label")
	btnMax := widget.NewButton("Max", func() {})
	btnMax.Importance = widget.HighImportance

	btns := [][]string{
		{"a", "America/Indiana/Indianapolis", "c", "d"},
		{"1", "America/North_Dakota/New_Salem", "3", "4"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"a", "b", "c", "d"},
		{"1", "2", "3", "4"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"a", "b", "c", "d"},
		{"1", "2", "3", "4"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"a", "b", "c", "d"},
		{"1", "2", "3", "4"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
		{"T", "K", "O", "K"},
	}

	btnsLayout := make([]fyne.CanvasObject, 0, 4)

	for i := 0; i < len(btns); i++ {
		for j := 0; j < len(btns[i]); j++ {
			label := widget.NewLabel(btns[i][j])
			label.Wrapping = fyne.TextWrapWord
			label.TextStyle = fyne.TextStyle{Bold: true}
			btnsLayout = append(btnsLayout, label)
		}
	}

	wrap := container.NewGridWrap(fyne.NewSize(160, 50), btnsLayout...)

	showLayout := container.NewBorder(tipsLabel, nil, nil, nil, container.NewScroll(wrap))
	//splitLayout := container.NewVSplit(container.NewVBox(leftLayout, rightLayout), showLayout)
	//splitLayout.SetOffset(0.1)

	//paddLayout := container.NewPadded(container.NewVBox(leftLayout, rightLayout), showLayout)

	buttom := widget.NewButton("btn", func() {})
	buttom.Importance = widget.HighImportance
	box := container.NewBorder(
		topLayout,
		buttom,
		nil, nil,
		container.NewBorder(container.NewVBox(leftLayout, rightLayout), nil, nil, nil, showLayout))

	w.SetContent(box)
	w.Resize(fyne.NewSize(1200, 650))
	w.ShowAndRun()
}
