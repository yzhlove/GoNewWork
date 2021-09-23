package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myapp := app.New()
	window := myapp.NewWindow("Server ToolBox")

	label := widget.NewLabel("Path:/Users/yostar/workSpace/GoNewWork")

	tips := dialog.NewInformation("server tool", "server box ... ", window)

	entry := widget.NewEntry()
	entry.MultiLine = true

	var addressStr string

	address := widget.NewRadioGroup([]string{
		"10.155.10.11/a",
		"10.155.10.12/b",
		"10.155.10.12/c",
	}, func(s string) {
		addressStr = s
		fmt.Println("address -> ", addressStr)
	})

	custome := dialog.NewCustom("select address", "ok", address, window)
	custome.SetOnClosed(func() {
		entry.Text = addressStr
		entry.Refresh()
	})

	help := container.NewHBox(label, layout.NewSpacer(), widget.NewToolbar(widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() { tips.Show() })))

	btnCSV := widget.NewButton("Gener CSV", func() {
		entry.Text = "generate csv file"
		entry.Refresh()
	})

	btnZIP := widget.NewButton("Update ZIP", func() {
		entry.Text = "update zip package"
		entry.Refresh()
		custome.Show()
	})

	btnGacha := widget.NewButton("Gacha Test", func() {
		entry.Text = "gacha test button"
		entry.Refresh()
	})

	btnMore := widget.NewButton("More", func() {
		entry.Text = "gacha test button"
		entry.Refresh()
	})

	btnBox := container.NewHBox(btnCSV, btnZIP, btnGacha, layout.NewSpacer(), btnMore)

	center := container.New(layout.NewVBoxLayout(), entry, layout.NewSpacer(), btnBox)

	show := container.New(layout.NewBorderLayout(help, nil, nil, nil), help, center)

	window.SetContent(show)
	window.Resize(fyne.NewSize(500, 300))
	window.ShowAndRun()
}
