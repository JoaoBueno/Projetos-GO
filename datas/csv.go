package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("tray_produtos_2022-08-16-13-25-56.csv")
	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = ';'
	reader.LazyQuotes = true

	records, err1 := reader.ReadAll()

	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println(records)
}
