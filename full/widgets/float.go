package widgets

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewFloatWidget(label string, placeholder string) fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(placeholder)
	errorLabel := widget.NewLabel("")
	entry.OnChanged = func(content string) {
		if _, err := strconv.ParseFloat(content, 64); err != nil {
			errorLabel.SetText("Digite um número decimal válido")
		} else {
			errorLabel.SetText("")
		}
		errorLabel.Refresh()
	}

	return container.NewVBox(widget.NewLabel(label), entry, errorLabel)
}
