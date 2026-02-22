package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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

func (c *Config) CreateMenuItems(window fyne.Window) {
	openMenuItem := fyne.NewMenuItem("Open...", func() {})
	saveMenuItem := fyne.NewMenuItem("Save", func() {})
	c.SaveMenuItem = saveMenuItem 
	c.SaveMenuItem.Disabled = true 
	saveAsMenuItem := fyne.NewMenuItem("Save as...", c.saveAsFunc(window))

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	window.SetMainMenu(menu)
}

func (c *Config) saveAsFunc(window fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error){
			if err != nil {
				dialog.ShowError(err, window)
			}
			if write == nil {
				// user cancelled
				return
			}

			write.Write([]byte(c.EditWidget.Text))
			c.CurrentFile = write.URI()
			defer write.Close()

			window.SetTitle(window.Title() + " - " + write.URI().Name())
			c.SaveMenuItem.Disabled = false
		}, window)
		saveDialog.Show()
	}
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
	myCfg.CreateMenuItems(window)

	// set content of the window
	window.SetContent(container.NewHSplit(myCfg.EditWidget, myCfg.PreviewWidget))

	// show and run the window
	window.Resize(fyne.Size{Width: 800, Height: 500})
	window.CenterOnScreen()
	window.ShowAndRun()

}