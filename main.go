package main

import (
	"io"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
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
	openMenuItem := fyne.NewMenuItem("Open...", c.openFunc(window))
	saveMenuItem := fyne.NewMenuItem("Save", c.saveFunc(window))
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

			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md") {
				dialog.ShowInformation("Error", "Please name your file with '.md' or '.MD' extension", window)
				return 
			}

			write.Write([]byte(c.EditWidget.Text))
			c.CurrentFile = write.URI()
			defer write.Close()

			window.SetTitle(window.Title() + " - " + write.URI().Name())
			c.SaveMenuItem.Disabled = false
		}, window)
		saveDialog.SetFileName("Untitled.MD")
		saveDialog.SetFilter(storage.NewExtensionFileFilter([]string{".md", ".MD"}))
		saveDialog.Show()
	}
}

func (c *Config) openFunc(window fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
			
			}

			if read == nil {
				// user clicked cancel
				return
			}
			defer read.Close()

			data, err := io.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			c.EditWidget.SetText(string(data))
			c.CurrentFile = read.URI()

			window.SetTitle(window.Title() + " - " + read.URI().Name())
			c.SaveMenuItem.Disabled = false
		}, window)
		openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".md", ".MD"}))
		openDialog.Show()
	}
}

func (c *Config) saveFunc(window fyne.Window) func() {
	return func() {
		if c.CurrentFile != nil {
			write, err := storage.Writer(c.CurrentFile)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			defer write.Close()
			write.Write([]byte(c.EditWidget.Text))
		}
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