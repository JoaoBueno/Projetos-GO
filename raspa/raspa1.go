package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "http://www.fazenda.df.gov.br/aplicacoes/legislacao/legislacao/TelaSaidaDocumento.cfm?txtNumero=82&txtAno=2018&txtTipo=7&txtParte=."

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Encontrar a tabela na página
	table := doc.Find("table")

	// Iterar por cada linha da tabela
	table.Find("tr").Each(func(i int, row *goquery.Selection) {
		// Extrair as colunas da linha
		cols := row.Find("td")

		// Verificar se a linha é uma linha válida
		if cols.Length() == 2 {
			// Extrair os valores das colunas
			empresa := strings.TrimSpace(cols.Eq(0).Text())
			empresa = strings.ReplaceAll(empresa, "\n", "")
			cnpj := strings.TrimSpace(cols.Eq(1).Text())
			cnpj = strings.ReplaceAll(cnpj, "\n", "")

			// Imprimir os valores
			fmt.Printf("%s: %s\n", empresa, cnpj)
		}
	})
}
