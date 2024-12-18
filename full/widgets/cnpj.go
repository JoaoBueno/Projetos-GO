package widgets

import (
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewCNPJWidget(label string, placeholder string) fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(placeholder)

	errorLabel := widget.NewLabel("")
	entry.OnChanged = func(content string) {
		if isValidCNPJ(content) {
			errorLabel.SetText("")
		} else {
			errorLabel.SetText("CNPJ inválido")
		}
		errorLabel.Refresh()
	}

	return container.NewVBox(widget.NewLabel(label), entry, errorLabel)
}

func isValidCNPJ(cnpj string) bool {
	re := regexp.MustCompile(`^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`)
	return re.MatchString(cnpj)
}
