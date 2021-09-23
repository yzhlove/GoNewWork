package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var ListItems = []string{"A", "B", "C", "A", "B", "C", "A", "B", "C", "A", "B", "C", "A", "B", "C", "A", "B", "C", "A", "B", "C"}

func ListView(event func(selected string)) *widget.List {
	list := widget.NewList(func() int {
		return len(ListItems)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("")
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		if ret, ok := object.(*widget.Label); ok {
			ret.SetText(ListItems[id])
			ret.Alignment = fyne.TextAlignCenter
			ret.Refresh()
		}
	})

	list.OnUnselected = func(id widget.ListItemID) {
		event(ListItems[id])
	}
	return list
}
