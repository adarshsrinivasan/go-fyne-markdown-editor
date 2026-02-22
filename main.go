package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
	button *widget.Button
	entry *widget.Entry
}

func (a *App) MakeUI() {
	a.output = widget.NewLabel("Hello, World!")
	a.entry = widget.NewEntry()
	a.button = widget.NewButton("Enter", func() {
		a.output.SetText(a.entry.Text)
	})
	a.button.Importance = widget.HighImportance
}

func main() {
    a := app.New()
	myApp := &App{}

	window := a.NewWindow("Hello World")
	myApp.MakeUI()

	window.SetContent(container.NewVBox(myApp.output, myApp.entry, myApp.button))
	window.Resize(fyne.Size{Width: 500, Height: 500})
	window.ShowAndRun()
}