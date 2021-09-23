package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"strconv"
	"time"
)

func main() {
	MyApp := app.New()
	c := MyApp.NewWindow("list 控件 列表")

	//lists(c)
	listObjs(c)

	c.Resize(fyne.NewSize(600, 600))
	c.ShowAndRun()

}

// list放入普通切片
func lists(w fyne.Window) {

	//data := []string{"red", "blue", "yellow"}  // 普通的切片

	data := findfont.List() // 字体库切片

	labelTatol := widget.NewLabel("")
	labelTatols := widget.NewLabel("")
	labelTatol.SetText(strconv.Itoa(len(data)))
	list1 := widget.NewList(
		func() int {
			return len(data) // 数据长度
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("") //标签容器
			//return widget.NewEntry()  //输入框容器
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			//fmt.Println(id, object)
			lable1 := object.(*widget.Label)
			lable1.Text = data[id]
		},
	)
	list1.Select(1) // 选中
	list1.OnUnselected = func(id widget.ListItemID) {
		fmt.Println(id, "---")
	}

	stopChan := make(chan struct{})
	// 开启自动预览
	buttonStart := widget.NewButton("start", func() {
		tickers := time.NewTicker(time.Millisecond * 200)
		go func() {
			for i := 0; i < len(data); i++ {
				fmt.Println(i, "--", len(data))
				select {
				case <-tickers.C:
					list1.Select(i)
					labelTatols.SetText(strconv.Itoa(i))
				case <-stopChan:
					return
				}
			}
		}()

	})
	// 停止自动预览
	buttonStop := widget.NewButton("stop", func() {
		stopChan <- struct{}{}
	})

	//w.SetContent(list1,labelTatol)
	w.SetContent(container.NewBorder(container.NewVBox(buttonStart, buttonStop), container.NewVBox(labelTatol, labelTatols), nil, nil, list1))
}

type person struct {
	name, age string
}

// list方入结构体
func listObjs(w fyne.Window) {

	data := []person{
		person{"zhangsan", "13"},
		person{"lisi", "24"},
		person{"wangwu", "35"},
	}

	list1 := widget.NewList(
		func() int {
			return len(data) // 数据长度
		},
		func() fyne.CanvasObject {
			//return container.NewHBox(widget.NewLabel(""), widget.NewLabel("")) //标签容器
			return widget.NewEntry() //输入框容器
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			//fmt.Println(id, object)

			entry, ok := object.(*widget.Entry)
			if !ok {
				panic("entry type error")
			}

			entry.Text = fmt.Sprintf("%s - %s", data[id].name, data[id].age)
			entry.Refresh()
			//lable1 := object.(*fyne.Container)
			//leftLabel := lable1.Objects[0].(*widget.Label)
			//rightLabel := lable1.Objects[1].(*widget.Label)
			//leftLabel.Text = data[id].name
			//rightLabel.Text = data[id].age
		},
	)
	list1.Select(1) // 选中
	list1.OnUnselected = func(id widget.ListItemID) {
		fmt.Println(id, "---")
	}

	w.SetContent(list1)
}
