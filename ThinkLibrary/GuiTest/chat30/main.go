package main

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
	"strings"
	"time"
)

func main() {
	//test1()
	test2()
}

func test2() {

	var count int
	MainWindow(func(win fyne.Window) fyne.CanvasObject {
		btn := widget.NewButton("OPEN", func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			count++
			log.Printf("OnTapped Event %d \n", count)
			progress := NewProgress(win, "progress")
			progress.Show()
			state := make(chan struct{})
			go func() {
				<-ctx.Done()
				progress.Hide()
				close(state)
			}()
			time.Sleep(time.Second * 2)
			cancel()
			<-state
			NewStd(win, "stander").Show()
		})

		return container.NewBorder(nil, btn, nil, nil)
	})
}

func test1() {

	var count int

	MainWindow(func(win fyne.Window) fyne.CanvasObject {

		btn := widget.NewButton("OPEN", func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			count++
			log.Printf("OnTapped Event %d \n", count)

			NewAsyncProgress(win, ctx, "progress frame widget")
			time.Sleep(time.Second * 2)
			cancel()

			NewStd(win, "stander frame widget").Show()
		})

		return container.NewBorder(nil, btn, nil, nil)
	})

}

func MainWindow(callback func(win fyne.Window) fyne.CanvasObject) {

	application := app.New()
	window := application.NewWindow("Test")

	window.SetContent(callback(window))
	window.Resize(fyne.NewSize(800, 500))
	window.ShowAndRun()
}

func NewProgress(w fyne.Window, s string) dialog.Dialog {
	ct := container.NewGridWrap(fyne.NewSize(400, 10), widget.NewProgressBarInfinite())
	return dialog.NewCustom(s, "close", ct, w)
}

func NewAsyncProgress(w fyne.Window, ctx context.Context, s string) {

	ct := container.NewGridWrap(fyne.NewSize(400, 10), widget.NewProgressBarInfinite())
	dg := dialog.NewCustom(s, "close", ct, w)
	go func() {
		defer dg.Hide()
		dg.Show()
		select {
		case <-ctx.Done():
			break
		}
	}()

}

func NewStd(w fyne.Window, s string) dialog.Dialog {
	labelText := widget.NewLabel(strings.Repeat("abc", 5))
	labelText.Wrapping = fyne.TextWrapBreak

	ct := container.NewGridWrap(fyne.NewSize(250, 100), container.NewVScroll(labelText))
	return dialog.NewCustom(s, "close", ct, w)
}
