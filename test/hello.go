package main

import (
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
    "time"
)


func main() {
	currentTime := time.Now()
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetFullScreen(true)
	w.SetContent(widget.NewVBox(
		widget.NewLabel(currentTime.Format("2006-01-02 3:4:5 pm")),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))
	
	w.ShowAndRun()

}
