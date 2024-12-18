package widgets

import (
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewEmailWidget(label string, placeholder string) fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(placeholder)

	errorLabel := widget.NewLabel("")
	entry.OnChanged = func(content string) {
		if isValidEmail(content) {
			errorLabel.SetText("")
		} else {
			errorLabel.SetText("E-mail inválido")
		}
		errorLabel.Refresh()
	}

	return container.NewVBox(widget.NewLabel(label), entry, errorLabel)
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
