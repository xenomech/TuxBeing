package main

import (
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
)
func main() {

	overidewindow := app.NewWindow("overide")
	overidewindow.SetContent(widget.NewVBox(
		widget.NewLabel("Are You Sure??"),
		widget.NewButton("Yes !",func(){overidewindow.Quit()}),
		widget.NewButton("No",func(){}),))
	overidewindow.ShowAndRun()
}