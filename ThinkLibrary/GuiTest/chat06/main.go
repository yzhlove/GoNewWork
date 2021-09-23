package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

type diagonal struct {
}

func (d *diagonal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

func (d *diagonal) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, containerSize.Height - d.MinSize(objects).Height)
	for _, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(size.Width, size.Height))
	}
}

//func main() {
//	a := app.New()
//	w := a.NewWindow("Diagonal")
//
//	text1 := widget.NewLabel("topleft")
//	text2 := widget.NewLabel("Middle Label")
//	text3 := widget.NewLabel("bottomright")
//
//	w.SetContent(container.New(&diagonal{}, text1, text2, text3))
//	w.ShowAndRun()
//}

func main() {
	a := app.NewWithID("com.example.tutorial.preferences")
	w := a.NewWindow("Timeout")

	var timeout time.Duration

	timeoutSelector := widget.NewSelect([]string{"10 seconds", "30 seconds", "1 minute"}, func(selected string) {
		switch selected {
		case "10 seconds":
			timeout = 10 * time.Second
		case "30 seconds":
			timeout = 30 * time.Second
		case "1 minute":
			timeout = time.Minute
		}

		a.Preferences().SetString("AppTimeout", selected)
	})

	timeoutSelector.SetSelected(a.Preferences().StringWithFallback("AppTimeout", "10 seconds"))

	go func() {
		time.Sleep(timeout)
		a.Quit()
	}()

	w.SetContent(timeoutSelector)
	w.ShowAndRun()
}