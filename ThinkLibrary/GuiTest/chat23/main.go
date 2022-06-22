package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strings"
)

func main() {

	rich := widget.NewRichText()
	rich.Wrapping = fyne.TextWrapBreak
	scroll := container.NewVScroll(rich)
	str1 := strings.Repeat("A", 100)
	str2 := strings.Repeat("B", 100)
	MainWindow(func(win fyne.Window) fyne.CanvasObject {
		rich.Segments = append(rich.Segments,
			&widget.TextSegment{
				Text: str1,
				//Style: widget.RichTextStyleBlockquote,
			})
		rich.Segments = append(rich.Segments,
			&widget.TextSegment{
				Text:  str2,
				Style: widget.RichTextStyleCodeBlock,
			})
		rich.Segments = append(rich.Segments,
			&widget.TextSegment{
				Text:  str1,
				Style: widget.RichTextStyleCodeInline,
			})
		rich.Segments = append(rich.Segments,
			&widget.TextSegment{
				Text:  str2,
				Style: widget.RichTextStyleEmphasis,
			})
		rich.Segments = append(rich.Segments,
			&widget.TextSegment{
				Text:  str1,
				Style: widget.RichTextStyleHeading,
			})
		rich.Segments = append(rich.Segments,
			&widget.TextSegment{
				Text:  str2,
				Style: widget.RichTextStyleInline,
			})
		rich.Segments = append(rich.Segments,
			&widget.TextSegment{
				Text:  str1,
				Style: widget.RichTextStyleParagraph,
			})
		rich.Segments = append(rich.Segments,
			&widget.TextSegment{
				Text:  str2,
				Style: widget.RichTextStylePassword,
			})
		rich.Segments = append(rich.Segments,
			&widget.TextSegment{
				Text:  str1,
				Style: widget.RichTextStyleStrong,
			})
		rich.Segments = append(rich.Segments,
			&widget.TextSegment{
				Text:  str2,
				Style: widget.RichTextStyleSubHeading,
			})
		return scroll
	})

}

func MainWindow(callback func(win fyne.Window) fyne.CanvasObject) {

	application := app.New()
	window := application.NewWindow("Test")

	window.SetContent(callback(window))
	window.Resize(fyne.NewSize(600, 400))
	window.ShowAndRun()
}
