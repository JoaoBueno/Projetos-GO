package main

import (
	"C"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)
import "unsafe"

func csvToMap(reader io.Reader, separador rune, lazyQuotes bool) ([]string, []map[string]string, error) {
	r := csv.NewReader(reader)
	r.Comma = separador
	r.LazyQuotes = lazyQuotes

	rows := []map[string]string{}
	var header []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			return header, rows, err
		}
		if err != nil {
			return header, rows, err
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
	return header, rows, nil
}

/*CSVLoad is a func to load a CSV file*/
//export CSVLoad
func CSVLoad(str *C.char, separador *C.char, lazyQuotes *C.char, header **C.char, csv **C.char, ret *int) {
	filename := C.GoString(str)
	sepa, _ := utf8.DecodeRune(C.GoString(separador)[0])
	lazyq := C.GoString(lazyQuotes)
	fmt.Println(filename, separador, lazyq)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		*ret = 10001 // erro ao abrir o arquivo
		return
	}

	h, m, err := csvToMap(bufio.NewReader(file), sepa, lazyq)

	if err != nil {
		fmt.Println(err)
		*ret = 10002 // erro ao ler o arquivo
		return
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(&dest[0])), len(src)), src)
	*header = C.CString(h)
	*csv = C.CString(m)
	*ret = 0
	return
}

func main() {}
