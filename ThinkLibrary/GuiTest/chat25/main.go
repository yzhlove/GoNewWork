package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const content = `
Knowledge is one thing, virtue is another; good sense is not conscience, 
refinement is not humility nor is largeness and justness of view faith. Philosophy, however enlightened, however profound,
gives no command over the passions, no influential motives, no vivifying principles. Liberal Education makes not the Christian, 
not the Catholic, but the gentleman. It is well to be a gentleman, it is well to have a cultivated intellect, a delicate taste, 
a candid, equitable, dispassionate mind, a noble and courteous bearing in the conduct of life—these are the connatural qualities of a large knowledge; 
they are the objects of a University. I am advocating, I shall illustrate and insist upon them; but still, I repeat, they 
are no guarantee for sanctity or even for conscientiousness, and they may attach to the man of the world, to the profligate
, to the heartless, pleasant, alas, and attractive as he shows when decked out in them. Taken by themselves, they do but seem
to be what they are not; they look like virtue at a distance, but they are detected by close observers, and in the long run; 
and hence it is that they are popularly accused of pretense and hypocrisy, not, I repeat, from their own fault, but because 
their professors and their admirers persist in taking them for what they are not, and are officious in arrogating for them 
a praise to which they have no claim. Quarry the granite rock with razors, or moor the vessel with a thread of silk, then may 
you hope with such keen and delicate instruments as human knowledge and human reason to contend against those giants, the passion and the pride of man.
`

func main() {

	MainWindow(func(win fyne.Window) fyne.CanvasObject {

		richText := widget.NewRichText()
		richText.Segments = append(richText.Segments, &widget.TextSegment{Text: content})
		richText.Segments = append(richText.Segments, &widget.TextSegment{Text: content})
		return container.NewScroll(richText)
	})

}

func MainWindow(callback func(win fyne.Window) fyne.CanvasObject) {

	application := app.New()
	window := application.NewWindow("Test")

	window.SetContent(callback(window))
	window.Resize(fyne.NewSize(300, 400))
	window.ShowAndRun()
}
