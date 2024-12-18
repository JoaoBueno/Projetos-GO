package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/therecipe/qt/widgets"
)

type Menu struct {
	widgets.QMainWindow

	_ func() `constructor:"init"`
}

func NewMenu() *Menu {
	return &Menu{}
}

func (m *Menu) init() {
	m.SetWindowTitle("CSV Menu Example")
	m.Resize2(600, 400)

	bar := m.MenuBar()

	csvFile, err := os.Open("1linha.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	m.createMenuLevel(bar, records, 0, "")
}

func (m *Menu) createMenuLevel(parent widgets.QWidget_ITF, records [][]string, level int, parentName string) {
	for _, record := range records {
		rowLevel, _ := strconv.Atoi(record[0])
		if rowLevel != level || (parentName != "" && !strings.HasPrefix(record[4], parentName)) {
			continue
		}

		menuText := strings.TrimPrefix(record[4], parentName)
		menuText = strings.Replace(menuText, "&", "", -1)

		if record[2] == " " {
			newMenu := widgets.NewQMenu2(menuText, nil)
			if menu, ok := parent.(*widgets.QMenu); ok {
				menu.AddMenu(newMenu)
				m.createMenuLevel(newMenu, records, level+1, record[4])
			} else if bar, ok := parent.(*widgets.QMenuBar); ok {
				bar.AddMenu(newMenu)
				m.createMenuLevel(newMenu, records, level+1, record[4])
			}
		} else {
			action := widgets.NewQAction(parent)
			action.SetText(menuText)
			action.ConnectTriggered(func(bool) {
				m.actionTriggered(record[5])
			})
			if menu, ok := parent.(*widgets.QMenu); ok {
				menu.AddAction(menuText)
			}
		}
	}
}

func (m *Menu) actionTriggered(prog string) {
	fmt.Println("Action triggered for program:", prog)
	// Aqui você pode adicionar lógica para executar o programa ou fazer outra ação
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	menu := NewMenu()
	menu.Show()

	widgets.QApplication_Exec()
}
