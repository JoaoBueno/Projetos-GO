package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewBooleanWidget(label string) fyne.CanvasObject {
	check := widget.NewCheck(label, nil)
	return container.NewVBox(check)
}
