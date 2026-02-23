package main_test

import (
	"testing"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"github.com/adarshsrinivasan/go-fyne-markdown-editor"
)

func TestConfig_MakeUI(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		content string
	}{
		{"empty", ""},
		{"simple", "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c main.Config
			testApp := test.NewApp()
			testWindow := testApp.NewWindow("Test Window")
			c.MakeUI()


			testWindow.SetContent(container.NewHSplit(c.EditWidget, c.PreviewWidget))

			testApp.Run()

			test.Type(c.EditWidget, tt.content)
			
			if tt.content != c.PreviewWidget.String() {
				t.Errorf("Preview widget text does not match expected content. Expected: %s, Got: %s", tt.content, c.PreviewWidget.String())
			}
		})
	}	
}
