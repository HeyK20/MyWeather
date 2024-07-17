package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	// Define progress bar and infinite progress bar
	progress := widget.NewProgressBar()
	infinite := widget.NewProgressBarInfinite()

	hello := widget.NewLabel("Hello Fyne!")
	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 1500)
			progress.SetValue(i)
		}
	}()
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!",
			func() {
				hello.SetText("Welcome :)")
			}),
		progress,
		infinite))

	w.ShowAndRun()
}
