package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ListView(event func(index widget.ListItemID)) *widget.List {
	list := widget.NewList(
		func() int { return len(ListItems) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.ListItemID, object fyne.CanvasObject) {
			if ret, ok := object.(*widget.Label); ok {
				ret.SetText(ListItems[id].Name)
				ret.Alignment = fyne.TextAlignLeading
				ret.Refresh()
			}
		})

	list.Select(0)
	list.OnSelected = func(id widget.ListItemID) {
		event(id)
	}

	return list
}
