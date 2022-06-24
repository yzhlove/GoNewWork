package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strings"
)

func TopLayout(path string) fyne.CanvasObject {
	root := widget.NewIcon(theme.HomeIcon())
	pathLabel := widget.NewLabel(path)
	pathLabel.Alignment = fyne.TextAlignLeading
	pathLabel.TextStyle = fyne.TextStyle{Bold: true}
	return container.NewHBox(root, pathLabel)
}

func LeftLayout(event func(id widget.ListItemID)) fyne.CanvasObject {
	return ListView(event)
}

func XlsxBox(event func(label *widget.Label)) fyne.CanvasObject {
	ret := widget.NewLabel("")
	scroll := container.NewScroll(ret)
	csvBtn := widget.NewButton("生成CSV文件", func() {
		event(ret)
	})
	csvBtn.Importance = widget.HighImportance
	return widget.NewCard("", "CSV文件生成工具", container.NewBorder(nil, csvBtn, nil, nil, scroll))
}

func UpdateZipBox(event func(label *widget.Label, s []string)) fyne.CanvasObject {

	var strs, tmp []string
	for _, k := range Address {
		strs = append(strs, fmt.Sprintf("%s-%s-%s", k.Name, k.Ip, k.Desc))
	}

	topLabel := widget.NewLabel("")
	hScroll := container.NewScroll(topLabel)

	group := widget.NewCheckGroup(strs, func(strings []string) {
		tmp = strings
	})
	vScroll := container.NewScroll(group)

	split := container.NewVSplit(hScroll, vScroll)
	split.SetOffset(0.4)

	utBtn := widget.NewButton("生成Zip包", func() {
		if len(tmp) > 0 {
			event(topLabel, tmp)
		}
	})
	utBtn.Importance = widget.HighImportance

	card := widget.NewCard("", "上传Update.zip包", container.NewBorder(nil, utBtn, nil, nil, split))
	card.Hide()
	return card
}

func GachaBox() fyne.CanvasObject {

	rich := widget.NewRichText()
	rich.Wrapping = fyne.TextWrapBreak
	scroll := container.NewVScroll(rich)
	str1 := strings.Repeat("你", 20) + "over"
	str2 := strings.Repeat("我", 25) + "over"
	rich.Segments = append(rich.Segments,
		&widget.TextSegment{
			Text: "1" + str1,
		})
	rich.Segments = append(rich.Segments,
		&widget.TextSegment{
			Text:  "2" + str2,
			Style: widget.RichTextStyleCodeBlock,
		})
	rich.Segments = append(rich.Segments,
		&widget.TextSegment{
			Text:  "3" + str1,
			Style: widget.RichTextStyleCodeInline,
		})
	rich.Segments = append(rich.Segments,
		&widget.TextSegment{
			Text:  "4" + str2,
			Style: widget.RichTextStyleEmphasis,
		})
	rich.Segments = append(rich.Segments,
		&widget.TextSegment{
			Text:  "5" + str1,
			Style: widget.RichTextStyleHeading,
		})
	rich.Segments = append(rich.Segments,
		&widget.TextSegment{
			Text:  "6" + str2,
			Style: widget.RichTextStyleInline,
		})
	rich.Segments = append(rich.Segments,
		&widget.TextSegment{
			Text:  "7" + str1,
			Style: widget.RichTextStyleParagraph,
		})
	rich.Segments = append(rich.Segments,
		&widget.TextSegment{
			Text:  "8" + str2,
			Style: widget.RichTextStylePassword,
		})
	rich.Segments = append(rich.Segments,
		&widget.TextSegment{
			Text:  "9" + str1,
			Style: widget.RichTextStyleStrong,
		})
	rich.Segments = append(rich.Segments,
		&widget.TextSegment{
			Text:  "10" + str2,
			Style: widget.RichTextStyleSubHeading,
		})

	return scroll
}

func BottomLayout() fyne.CanvasObject {
	ret := widget.NewIcon(resourceYostarlogoBlack1200x1200Png)
	return container.NewHBox(layout.NewSpacer(), ret)
}
