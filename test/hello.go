package main

import (
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
)

func main() {

	currentTime := time.Now()
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetFullScreen(true)
	w.SetContent(widget.NewVBox(
		widget.NewLabel(currentTime.Format("2006-01-02 3:4:5 pm")),
		canvas.Refresh(widget.NewLabel),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))
	x.image.File = file.Name()
	x.image = &canvas.Image{FillMode: canvas.ImageFillOriginal}

	w.ShowAndRun()

}
