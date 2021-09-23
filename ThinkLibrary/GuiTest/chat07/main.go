package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"time"
)

func main1() {

	display := widget.NewEntry()
	display.MultiLine = true

	digits := []string{
		"7", "8", "9", "×",
		"4", "5", "6", "-",
		"1", "2", "3", "+",
	}
	var digitBtns []fyne.CanvasObject
	for _, val := range digits {
		digitBtns = append(digitBtns, widget.NewButton(val, nil))
	}
	digitContainer := fyne.NewContainerWithLayout(
		layout.NewGridLayout(4),
		digitBtns...)

	clearBtn := widget.NewButton("AC", nil)
	signBtn := widget.NewButton("+/-", nil)
	percentBtn := widget.NewButton("%", nil)
	divideBtn := widget.NewButton("÷", nil)
	clearContainer := fyne.NewContainerWithLayout(
		layout.NewGridLayout(4),
		clearBtn,
		signBtn,
		percentBtn,
		divideBtn,
	)

	zeroBtn := widget.NewButton("0", nil)
	dotBtn := widget.NewButton(".", nil)
	equalBtn := widget.NewButton("=", nil)
	zeroContainer := fyne.NewContainerWithLayout(
		layout.NewGridLayout(2),
		zeroBtn,
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(2),
			dotBtn,
			equalBtn,
		),
	)

	container := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		display,
		clearContainer,
		digitContainer,
		zeroContainer,
		nil,
	)

	w := app.New()
	ww := w.NewWindow("aaa")
	ww.SetContent(container)

	ww.ShowAndRun()
}



func label() {

	myApp := app.New()
	myWin := myApp.NewWindow("Test Label")

	l1 := widget.NewLabel("Name")
	l2 := widget.NewLabel("da\njun")

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), l1, l2)
	myWin.SetContent(container)
	myWin.Resize(fyne.NewSize(150, 150))
	myWin.ShowAndRun()

}

func btn() {
	myApp := app.New()
	myWin := myApp.NewWindow("Button")

	btn1 := widget.NewButton("text button", func() {
		fmt.Println("text button clicked")
	})

	btn2 := widget.NewButtonWithIcon("icon", theme.HomeIcon(), func() {
		fmt.Println("icon button clicked")
	})

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), btn1, btn2)
	myWin.SetContent(container)
	myWin.Resize(fyne.NewSize(150, 50))
	myWin.ShowAndRun()
}

func box() {

	myApp := app.New()
	myWin := myApp.NewWindow("Box")

	bbox := layout.NewVBoxLayout()
	bbox.Layout([]fyne.CanvasObject{
		widget.NewLabel("The top row of VBox"),
		widget.NewLabel("Label 1"),
		widget.NewLabel("Label 2"),
	}, fyne.NewSize(100, 30))

	container := fyne.NewContainerWithLayout(bbox)

	myWin.SetContent(container)
	myWin.Resize(fyne.NewSize(150, 150))
	myWin.ShowAndRun()

}

func layoutA() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text2.Move(fyne.NewPos(20, 20))
	//content := container.NewWithoutLayout(text1, text2)
	content := container.New(layout.NewGridLayout(2), text1, text2)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func boxLayout() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")

	text1 := canvas.NewText("Hello", color.Black)
	text2 := canvas.NewText("There", color.Black)
	text3 := canvas.NewText("(right)", color.Black)
	content := container.New(layout.NewHBoxLayout(), text1, layout.NewSpacer(), text2, text3)

	text4 := canvas.NewText("centered", color.Black)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, layout.NewSpacer(), centered))
	myWindow.ShowAndRun()
}

func login() {

	myApp := app.New()
	myWin := myApp.NewWindow("Entry")
	myWin.Resize(fyne.NewSize(300, 400))

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("input name")

	nameEntry.OnChanged = func(content string) {
		fmt.Println("name:", nameEntry.Text, "entered")
	}

	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("input password")

	nameBox := container.New(layout.NewHBoxLayout(), widget.NewLabel("Name"), nameEntry)
	passwoedBox := container.New(layout.NewHBoxLayout(), widget.NewLabel("Password"), passEntry)

	loginBtn := widget.NewButton("Login", func() {
		fmt.Println("name:", nameEntry.Text, "password:", passEntry.Text, "login in")
	})

	multiEntry := widget.NewEntry()
	multiEntry.SetPlaceHolder("please enter\nyour description")
	multiEntry.MultiLine = true

	content := container.New(layout.NewVBoxLayout(), nameBox, passwoedBox, loginBtn, multiEntry)
	myWin.SetContent(content)
	myWin.ShowAndRun()

}

func register() {
	myApp := app.New()
	myWin := myApp.NewWindow("Choices")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("input name")

	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("input password")

	repeatPassEntry := widget.NewPasswordEntry()
	repeatPassEntry.SetPlaceHolder("repeat password")

	nameBox := container.New(layout.NewHBoxLayout(), widget.NewLabel("Name"), layout.NewSpacer(), nameEntry)
	passwordBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), passEntry)
	repeatPasswordBox := container.New(layout.NewHBoxLayout(), widget.NewLabel("Repeat Password"), layout.NewSpacer(), repeatPassEntry)

	sexRadio := widget.NewRadioGroup([]string{"male", "female", "unknown"}, func(value string) {
		fmt.Println("sex:", value)
	})

	sexBox := container.New(layout.NewHBoxLayout(), widget.NewLabel("Sex"), sexRadio)

	football := widget.NewCheck("football", func(value bool) {
		fmt.Println("football:", value)
	})
	basketball := widget.NewCheck("basketball", func(value bool) {
		fmt.Println("basketball:", value)
	})
	pingpong := widget.NewCheck("pingpong", func(value bool) {
		fmt.Println("pingpong:", value)
	})
	hobbyBox := container.New(layout.NewHBoxLayout(), widget.NewLabel("Hobby"), football, basketball, pingpong)

	provinceSelect := widget.NewSelect([]string{"anhui", "zhejiang", "shanghai"}, func(value string) {
		fmt.Println("province:", value)
	})
	provinceBox := container.New(layout.NewHBoxLayout(), widget.NewLabel("Province"), layout.NewSpacer(), provinceSelect)

	registerBtn := widget.NewButton("Register", func() {
		fmt.Println("name:", nameEntry.Text, "password:", passEntry.Text, "register")
	})

	content := container.New(layout.NewVBoxLayout(), nameBox, passwordBox, repeatPasswordBox,
		sexBox, hobbyBox, provinceBox, registerBtn)
	myWin.SetContent(content)
	myWin.ShowAndRun()
}

func from() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form")

	nameEntry := widget.NewEntry()
	passEntry := widget.NewPasswordEntry()

	form := widget.NewForm(
		&widget.FormItem{Text: "Name", Widget: nameEntry},
		&widget.FormItem{Text: "Pass", Widget: passEntry},
	)
	form.OnSubmit = func() {
		fmt.Println("name:", nameEntry.Text, "pass:", passEntry.Text, "login in")
	}
	form.OnCancel = func() {
		fmt.Println("login canceled")
	}

	myWindow.SetContent(form)
	myWindow.Resize(fyne.NewSize(150, 150))
	myWindow.ShowAndRun()
}

func progress() {

	myApp := app.New()
	myWindow := myApp.NewWindow("ProgressBar")

	bar1 := widget.NewProgressBar()
	bar1.Min = 0
	bar1.Max = 100
	bar2 := widget.NewProgressBarInfinite()

	go func() {
		for i := 0; i <= 100; i++ {
			time.Sleep(time.Millisecond * 500)
			bar1.SetValue(float64(i))
		}
	}()

	content := container.New(layout.NewVBoxLayout(), bar1, bar2)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(150, 150))
	myWindow.ShowAndRun()

}

func tableSelect() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer")

	nameLabel := widget.NewLabel("Name: dajun")
	sexLabel := widget.NewLabel("Sex: male")
	ageLabel := widget.NewLabel("Age: 18")
	addressLabel := widget.NewLabel("Province: shanghai")
	addressLabel.Hide()
	profile := container.New(layout.NewVBoxLayout(), nameLabel, sexLabel, ageLabel, addressLabel)

	musicRadio := widget.NewRadioGroup([]string{"on", "off"}, func(string) {})
	showAddressCheck := widget.NewCheck("show address?", func(value bool) {
		if !value {
			addressLabel.Hide()
		} else {
			addressLabel.Show()
		}
	})
	memberTypeSelect := widget.NewSelect([]string{"junior", "senior", "admin"}, func(string) {})

	setting := widget.NewForm(
		&widget.FormItem{Text: "music", Widget: musicRadio},
		&widget.FormItem{Text: "check", Widget: showAddressCheck},
		&widget.FormItem{Text: "member type", Widget: memberTypeSelect},
	)

	tabs := container.NewAppTabs(container.NewTabItem("Profile", profile), container.NewTabItem("Setting", setting))

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}

func boxHV() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")

	hcontainer1 := container.New(layout.NewHBoxLayout(), canvas.NewText("left", color.Black), canvas.NewText("right", color.Black))

	// 左对齐
	hcontainer2 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), canvas.NewText("left", color.Black), canvas.NewText("right", color.Black))

	// 右对齐
	hcontainer3 := container.New(layout.NewHBoxLayout(), canvas.NewText("left", color.Black), canvas.NewText("right", color.Black), layout.NewSpacer())

	// 中间对齐
	hcontainer4 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), canvas.NewText("left", color.Black), canvas.NewText("right", color.Black),
		layout.NewSpacer())

	// 两边对齐
	hcontainer5 := container.New(layout.NewHBoxLayout(),
		canvas.NewText("left", color.Black),
		layout.NewSpacer(),
		canvas.NewText("right", color.Black))

	myWindow.SetContent(container.New(layout.NewVBoxLayout(),
		hcontainer1, hcontainer2, hcontainer3, hcontainer4, hcontainer5))
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}

func gridLayout() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	img1 := canvas.NewImageFromResource(theme.FileIcon())
	img2 := canvas.NewImageFromResource(theme.FileIcon())
	img3 := canvas.NewImageFromResource(theme.FileIcon())
	myWindow.SetContent(container.New(layout.NewGridLayout(2),
		img1, img2, img3))
	myWindow.Resize(fyne.NewSize(300, 300))
	myWindow.ShowAndRun()
}

func gridWraplayout() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Wrap Layout")

	img1 := canvas.NewImageFromResource(theme.FyneLogo())
	img2 := canvas.NewImageFromResource(theme.FyneLogo())
	img3 := canvas.NewImageFromResource(theme.FyneLogo())
	myWindow.SetContent(
		container.New(
			layout.NewGridWrapLayout(fyne.NewSize(150, 150)),
			img1, img2, img3))
	myWindow.ShowAndRun()
}

func toolbar() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Toolbar")

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			fmt.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {
			fmt.Println("Cut")
		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			fmt.Println("Copy")
		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			fmt.Println("Paste")
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	content := container.New(
		layout.NewBorderLayout(toolbar, nil, nil, nil),
		toolbar, widget.NewLabel(`Lorem ipsum dolor, 
    sit amet consectetur adipisicing elit.
    Quidem consectetur ipsam nesciunt,
    quasi sint expedita minus aut,
    porro iusto magnam ducimus voluptates cum vitae.
    Vero adipisci earum iure consequatur quidem.`),
	)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func borderLayout() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	left := canvas.NewText("left", color.Black)
	right := canvas.NewText("right", color.Black)
	top := canvas.NewText("top", color.Black)
	bottom := canvas.NewText("bottom", color.Black)
	content := widget.NewLabel(`Lorem ipsum dolor, 
  sit amet consectetur adipisicing elit.
  Quidem consectetur ipsam nesciunt,
  quasi sint expedita minus aut,
  porro iusto magnam ducimus voluptates cum vitae.
  Vero adipisci earum iure consequatur quidem.`)

	container := container.New(
		layout.NewBorderLayout(top, bottom, left, right),
		top, bottom, left, right, content,
	)
	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}




func main() {

	//btn()
	//box()
	//entry()

	//layoutA()
	//boxLayout()

	//login()
	//register()
	//from()
	//progress()
	//tableSelect()

	//boxHV()
	//gridLayout()
	//gridWraplayout()
	//toolbar()
	borderLayout()
}