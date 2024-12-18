package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func CSVToMap(reader io.Reader) ([]string, []map[string]string) {
	r := csv.NewReader(reader)
	r.Comma = ';'
	r.LazyQuotes = true

	rows := []map[string]string{}
	var header []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header == nil {
			// if record[0][0] == 239 && record[0][1] == 187 && record[0][2] == 191 {
			// 	fmt.Println(record[0][3:])
			// }
			header = record
		} else {
			dict := map[string]string{}

			for i := range header {
				dict[header[i]] = record[i]
				// fmt.Println(header[i], record[i])
			}
			rows = append(rows, dict)
		}
	}
	return header, rows
}

func main() {
	file, err := os.Open("tray_teste.csv")
	if err != nil {
		fmt.Println(err)
	}

	h, m := CSVToMap(bufio.NewReader(file))

	// for a, row := range m {
	// 	for b, col := range row {
	// 		fmt.Printf("%d, %s: %s\n", a, b, col)
	// 	}
	// 	// fmt.Println(a, row["Código"], row["Altura do Produto"], row["Código Pai"], row["Descrição"])
	// }

	// fmt.Println(h)
	fmt.Println(m[0]["Código"])
	// fmt.Println(m[6]["ID"])

	filen, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(filen)
	writer.Comma = ';'
	// writer.LazyQuotes = true
	defer writer.Flush()

	err = writer.Write(h)
	checkError("Cannot write to file", err)

	linha := []string{}
	for _, row := range m {
		for _, head := range h {
			// fmt.Printf("%s -> %s\n", head, row[head])
			linha = append(linha, row[head])
			// fmt.Println(linha)
		}
		err = writer.Write(linha)
		checkError("Cannot write to file", err)
		linha = []string{}
	}

	// for _, value := range m {
	// 	fmt.Printf("value = %T\n", value)
	// 	// err := writer.Write(value)
	// 	// checkError("Cannot write to file", err)
	// }
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
