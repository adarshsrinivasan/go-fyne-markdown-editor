package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
    a := app.New()
	window := a.NewWindow("Hello World")
	label := widget.NewLabel("Hello, Fyne!")
	window.SetContent( container.NewVBox(
		label,
		widget.NewButton("Click me", func(){
			label.SetText("Button clicked")
		}),
	))
	window.ShowAndRun()
}