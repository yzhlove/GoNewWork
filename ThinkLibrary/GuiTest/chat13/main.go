package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	c := app.NewWithID("yzhlove")
	w := c.NewWindow("timeZoneTool")

	top := topLayout("what are you doing?")
	left := leftLayout()

	var items = []listViewItem{{name: "time_zone", canvas: left}}

	ctrl := listView(items, func(idx widget.ListItemID) {
		fmt.Println("=======> idx ==> ", idx)
	})

	center := container.NewHSplit(ctrl, left)
	center.SetOffset(0.3)

	w.SetContent(container.NewBorder(
		top, nil, nil, nil, center))
	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()
}

func topLayout(path string) fyne.CanvasObject {
	home := widget.NewIcon(theme.HomeIcon())
	descLabe := widget.NewLabel(path)
	descLabe.Alignment = fyne.TextAlignLeading
	descLabe.TextStyle = fyne.TextStyle{Italic: true, Bold: true}
	return container.NewHBox(home, descLabe)
}

func leftLayout() fyne.CanvasObject {

	itLabe := widget.NewLabel("Input Time:")
	itEntr := widget.NewEntry()
	itEntr.OnChanged = func(s string) {
		fmt.Println("input time:" + s)
	}
	//oneLayout := container.NewHBox(itLabe, )
	oneLayout := container.NewBorder(nil, nil, itLabe, nil, container.NewPadded(itEntr))

	zoneInLabe := widget.NewLabel("Please Input Zone:")
	zoneInSelect := widget.NewSelect([]string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	}, func(s string) {
		fmt.Println("input select => " + s)
	})
	twoLayout := container.NewHBox(zoneInLabe, zoneInSelect)

	zoneOutLabe := widget.NewLabel("Please Output Zone:")
	zoneOutSelect := widget.NewSelect([]string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	}, func(s string) {
		fmt.Println("output select => " + s)
	})
	threeLayout := container.NewHBox(zoneOutLabe, zoneOutSelect)

	otLabe := widget.NewLabel("Output Time:")
	otEntr := widget.NewEntry()
	fourLayout := container.NewHBox(otLabe, otEntr)

	return container.NewVBox(oneLayout, twoLayout, threeLayout, fourLayout)
}

func listView(views []listViewItem, fn func(idx widget.ListItemID)) *widget.List {
	view := widget.NewList(
		func() int {
			return len(views)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		}, func(id widget.ListItemID, object fyne.CanvasObject) {
			if r, ok := object.(*widget.Label); ok {
				r.SetText(views[id].name)
				r.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
				r.Refresh()
			}
		})

	view.Select(0)
	view.OnSelected = fn

	return view
}

type listViewItem struct {
	name   string
	canvas fyne.CanvasObject
}
