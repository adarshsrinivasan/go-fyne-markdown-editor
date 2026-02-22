package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	EditWidget *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile fyne.URI
	SaveMenuItem *fyne.MenuItem
}

func (c *Config) MakeUI() {
	c.EditWidget = widget.NewMultiLineEntry()
	c.PreviewWidget = widget.NewRichTextFromMarkdown("")
	
	c.EditWidget.OnChanged = c.PreviewWidget.ParseMarkdown
}


func main() {
	// create the app object
	a := app.New()

	// create the config object
	myCfg := &Config{}

	// create the window
	window := a.NewWindow("Markdown")

	// get the user interface
	myCfg.MakeUI()

	// set content of the window
	window.SetContent(container.NewHSplit(myCfg.EditWidget, myCfg.PreviewWidget))

	// show and run the window
	window.Resize(fyne.Size{Width: 800, Height: 500})
	window.CenterOnScreen()
	window.ShowAndRun()

}