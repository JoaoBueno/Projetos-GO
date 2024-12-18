package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("tray_produtos_2022-08-16-13-25-56.csv")

	if err != nil {

		log.Fatal(err)
	}

	r := csv.NewReader(f)
	r.Comma = ';'
	r.LazyQuotes = true

	for {

		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		for value := range record {
			fmt.Printf("%s\n", record[value])
		}
	}
}
