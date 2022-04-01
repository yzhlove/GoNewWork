package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strings"
	"think-library/GuiTest/chat14/assest"
	"time"
)

func main() {

	c := app.NewWithID("12138")
	w := c.NewWindow("ToolBOX")

	var entryText strings.Builder

	top := zoneComponentByTop(func(s string) {
		entryText.WriteString(s)
	})
	centerIn := zoneComponentByCenter(func(s string) {
		entryText.WriteString(s)
	})
	centerOut := zoneComponentByCenter(func(s string) {
		entryText.WriteString(s)
	})
	res := zoneComponentByRes(func() string {
		temp := entryText.String()
		entryText.Reset()
		return temp
	})

	w.SetContent(container.NewBorder(
		top,
		res,
		nil,
		nil,
		container.NewGridWithColumns(2, centerIn, centerOut),
	))

	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()
}

func zoneComponentByTop(changed func(s string)) fyne.CanvasObject {

	inputLabel := widget.NewLabel("Input:")
	inputLabel.TextStyle = fyne.TextStyle{Italic: true, Bold: true}

	inputEntry := widget.NewEntry()
	inputEntry.OnChanged = changed
	inputEntry.TextStyle = fyne.TextStyle{Italic: true, Bold: true}
	inputEntry.PlaceHolder = "please input time , style for '2006-01-02 15:04:06'"
	inputEntry.Validator = func(s string) error {
		if len(s) > 0 {
			_, err := time.Parse("2006-01-02 15:04:05", s)
			return err
		}
		return nil
	}
	return container.NewBorder(
		nil,
		nil,
		inputLabel,
		nil,
		container.NewPadded(inputEntry),
	)
}

func zoneComponentByCenter(change func(s string)) fyne.CanvasObject {

	descLabel := widget.NewLabel("SelectZone")
	descLabel.TextStyle = fyne.TextStyle{Italic: true, Bold: true}

	zoneSelect := widget.NewSelect(assest.GetZones(), change)
	zoneSelect.Alignment = fyne.TextAlignLeading
	zoneSelect.SetSelectedIndex(0)

	return container.NewVBox(container.NewPadded(descLabel),
		container.NewPadded(zoneSelect))

}

func zoneComponentByRes(tapped func() string) fyne.CanvasObject {

	entry := widget.NewEntry()
	entry.TextStyle = fyne.TextStyle{Italic: true, Bold: true}

	btn := widget.NewButtonWithIcon("Transform", theme.ComputerIcon(), func() {
		entry.SetText(tapped())
	})
	btn.Importance = widget.HighImportance

	return container.NewBorder(
		nil,
		nil,
		nil,
		btn,
		entry)
}
