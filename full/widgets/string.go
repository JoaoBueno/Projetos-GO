package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Validação para strings
type StringValidation struct {
	MinLength int
	MaxLength int
}

func NewStringWidget(label string, placeholder string, validation StringValidation) fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(placeholder)
	errorLabel := widget.NewLabel("")

	entry.OnChanged = func(content string) {
		if validation.MinLength > 0 && len(content) < validation.MinLength {
			errorLabel.SetText("Comprimento mínimo: " + string(validation.MinLength))
		} else if validation.MaxLength > 0 && len(content) > validation.MaxLength {
			errorLabel.SetText("Comprimento máximo: " + string(validation.MaxLength))
		} else {
			errorLabel.SetText("")
		}
		errorLabel.Refresh()
	}

	return container.NewVBox(widget.NewLabel(label), entry, errorLabel)
}
