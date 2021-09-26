package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("test GUI")
	ListView(w)
	w.Resize(fyne.NewSize(1000, 800))
	w.ShowAndRun()
}

func ListView(w fyne.Window) {

	strs := []string{"apple", "balance", "center", "debug", "enough", "four"}

	var ls *widget.List

	bind := func(s string, opt int) {
		fmt.Println("s -> ", s, "opt -> ", opt)
		if opt == 0 && ls != nil {
			strs = strs[:len(strs)-1]
			ls.Refresh()
		}
	}

	ls = widget.NewList(
		func() int { return len(strs) },
		func() fyne.CanvasObject {
			label := widget.NewLabel("")
			btnOk := widget.NewButtonWithIcon("Ok", theme.ConfirmIcon(), nil)
			btnOk.Importance = widget.HighImportance
			btnDel := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), nil)
			btnDel.Importance = widget.HighImportance
			return container.NewHBox(label, layout.NewSpacer(), btnOk, btnDel)
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			box := object.(*fyne.Container)
			label := box.Objects[0].(*widget.Label)
			btnOk := box.Objects[2].(*widget.Button)
			btnDel := box.Objects[3].(*widget.Button)

			if strs[id] == "apple" {
				btnDel.Importance = widget.MediumImportance
				btnDel.Refresh()
			}
			label.SetText(strs[id])
			btnOk.OnTapped = func() { bind(strs[id], 1) }
			btnDel.OnTapped = func() { bind(strs[id], 0) }
		})

	w.SetContent(ls)
}
