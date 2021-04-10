package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Calc")
	window.SetContent(
		widget.NewLabel("Hello"),
	)
	window.ShowAndRun()
}
