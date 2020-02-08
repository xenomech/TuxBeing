package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	overidewindow := app.NewWindow("overide")
	overidewindow.SetContent(widget.NewVBox(
		widget.NewLabel("Are You Sure??"),
		widget.NewButton("Yes !", func() { app.Quit() }),
		widget.NewButton("No", func() {})))
	overidewindow.ShowAndRun()
}
