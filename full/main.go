package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"full/widgets"
)

type Field struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Label       string `json:"label"`
	Placeholder string `json:"placeholder"`
}

type Table struct {
	Name   string  `json:"name"`
	Label  string  `json:"label"`
	Fields []Field `json:"fields"`
}

type Config struct {
	Tables []Table `json:"tables"`
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func generateField(field Field, w fyne.Window) fyne.CanvasObject {
	switch field.Type {
	case "string":
		entry := widget.NewEntry()
		entry.SetPlaceHolder(field.Label)
		return container.NewVBox(widget.NewLabel(field.Label), entry)

	case "integer":
		entry := widget.NewEntry()
		entry.SetPlaceHolder(field.Label)
		errorLabel := widget.NewLabel("")
		entry.OnChanged = func(content string) {
			if _, err := strconv.Atoi(content); err != nil {
				errorLabel.SetText("Digite apenas números inteiros")
			} else {
				errorLabel.SetText("")
			}
			errorLabel.Refresh()
		}
		return container.NewVBox(widget.NewLabel(field.Label), entry, errorLabel)

	case "email":
		return widgets.NewEmailWidget(field.Label, field.Placeholder)

	case "date":
		return widgets.NewDateWidget(field.Label, field.Placeholder, w)

	case "cnpj":
		return widgets.NewCNPJWidget(field.Label, field.Placeholder)

	default:
		return container.NewVBox(widget.NewLabel("Tipo de campo desconhecido"), widget.NewLabel(field.Label))
	}
}

func generateForm(table Table, w fyne.Window) fyne.CanvasObject {
	formItems := []fyne.CanvasObject{}

	for _, field := range table.Fields {
		formItems = append(formItems, generateField(field, w))
	}

	saveButton := widget.NewButton("Salvar", func() {
		fmt.Println("Dados salvos!")
	})
	formItems = append(formItems, saveButton)

	return container.NewVBox(formItems...)
}

func main() {
	app := app.New()
	w := app.NewWindow("Dynamic Form Builder")

	config, err := loadConfig("form.json")
	if err != nil {
		dialog.ShowError(err, w)
		log.Fatal(err)
	}

	if len(config.Tables) > 0 {
		form := generateForm(config.Tables[0], w)
		w.SetContent(form)
	} else {
		dialog.ShowInformation("Erro", "Nenhuma tabela definida", w)
	}

	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
