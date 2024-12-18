package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewDateWidget(label string, placeholder string, w fyne.Window) fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(placeholder)
	// entry.OnTapped = func() {
	// 	dialog.ShowDatePicker("Selecione a data", time.Now(), func(date time.Time) {
	// 		entry.SetText(date.Format("02/01/2006"))
	// 	}, w)
	// }
	return container.NewVBox(widget.NewLabel(label), entry)
}
