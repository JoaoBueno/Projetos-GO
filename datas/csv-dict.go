package main

import (
	"fmt"
	"io"
	"os"

	csv "github.com/whosonfirst/go-whosonfirst-csv"
)

func main() {
	reader, err := csv.NewDictReaderFromPath("tray_teste.csv", ';', true)
	// reader.Comma = ';'
	// reader.LazyQuotes = true

	if err != nil {
		panic(err)
	}

	fmt.Println(reader)

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(row["Código"])

		// value, ok := row["Código"]

		// and so on...
	}

	writer, err := csv.NewDictWriter(os.Stdout, fieldnames)

	// for a, row := range m {
	// 	fmt.Println(a, row["Código"], row["Altura do Produto"], row["Código Pai"])
	// }

}
