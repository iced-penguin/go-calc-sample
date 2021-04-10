package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Calc")
	window.SetContent(
		widget.NewLabel("Hello"),
	)

	app.Settings().SetTheme(theme.DarkTheme())
	window.ShowAndRun()
}
