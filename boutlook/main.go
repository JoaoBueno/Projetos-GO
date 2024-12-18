package main

import (
	"encoding/json"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// FieldConfig representa a estrutura do campo no JSON
type FieldConfig struct {
	Label  string  `json:"label"`
	PosX   float32 `json:"posX"`
	PosY   float32 `json:"posY"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
}

// TelaConfig representa a configuração da tela
type TelaConfig struct {
	Fields []FieldConfig `json:"fields"`
}

// CarregarJSON carrega o conteúdo do arquivo JSON
func CarregarJSON(nomeArquivo string) (*TelaConfig, error) {
	file, err := os.Open(nomeArquivo)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config TelaConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Telas Dinâmicas com JSON")

	// Carregar configuração do JSON
	config, err := CarregarJSON("tela_config.json")
	if err != nil {
		fmt.Println("Erro ao carregar JSON:", err)
		return
	}

	// Criar widgets com base no JSON
	var widgets []fyne.CanvasObject
	for _, field := range config.Fields {
		// Rótulo do campo
		label := widget.NewLabel(field.Label)
		label.Move(fyne.NewPos(field.PosX, field.PosY))
		label.Resize(fyne.NewSize(field.Width, 30))

		// Campo de entrada
		entry := widget.NewEntry()
		entry.Move(fyne.NewPos(field.PosX, field.PosY+27))
		entry.Resize(fyne.NewSize(field.Width, field.Height))

		// Adicionar ao array de widgets
		widgets = append(widgets, label, entry)
	}

	// Criar container com layout livre (absoluto)
	content := container.NewWithoutLayout(widgets...)
	myWindow.SetContent(content)

	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
