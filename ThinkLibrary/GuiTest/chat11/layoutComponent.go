package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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

func GachaBox(event func(label *widget.Label)) fyne.CanvasObject {
	ret := widget.NewLabel("")
	scroll := container.NewScroll(ret)

	btn := widget.NewButton("查看Gacha测试结果", func() { event(ret) })
	btn.Importance = widget.HighImportance

	card := widget.NewCard("", "Gacha测试用例结果", container.NewBorder(nil, btn, nil, nil, scroll))
	card.Hide()
	return card
}

func BottomLayout() fyne.CanvasObject {
	ret := widget.NewIcon(resourceYostarlogoBlack1200x1200Png)
	return container.NewHBox(layout.NewSpacer(), ret)
}
