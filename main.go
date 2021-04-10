package main

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Calc")

	formula := ""
	label := widget.NewLabel(formula)
	label.Alignment = fyne.TextAlignTrailing

	pushNum := func(n int) {
		// 数式の先頭が0とならないように
		if formula == "" && n == 0 {
			return
		}
		formula += strconv.Itoa(n)
		label.SetText(formula)
	}

	clear := func() {
		formula = ""
		label.SetText(formula)
	}

	window.SetContent(
		widget.NewVBox(
			label,
			widget.NewHBox(
				fyne.NewContainerWithLayout(
					layout.NewFixedGridLayout(fyne.NewSize(200, 200)),
					fyne.NewContainerWithLayout(
						layout.NewGridLayout(3),
						widget.NewButton("7", func() { pushNum(7) }),
						widget.NewButton("8", func() { pushNum(8) }),
						widget.NewButton("9", func() { pushNum(9) }),
						widget.NewButton("4", func() { pushNum(4) }),
						widget.NewButton("5", func() { pushNum(5) }),
						widget.NewButton("6", func() { pushNum(6) }),
						widget.NewButton("1", func() { pushNum(1) }),
						widget.NewButton("2", func() { pushNum(2) }),
						widget.NewButton("3", func() { pushNum(3) }),
						widget.NewButton("0", func() { pushNum(0) }),
						layout.NewSpacer(),
						widget.NewButton("=", nil),
					),
				),
				widget.NewVBox(
					widget.NewButton("CL", clear),
					widget.NewButton("/", nil),
					widget.NewButton("*", nil),
					widget.NewButton("+", nil),
					widget.NewButton("-", nil),
				),
			),
		),
	)

	app.Settings().SetTheme(theme.DarkTheme())
	window.ShowAndRun()
}
