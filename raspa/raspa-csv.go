package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "http://www.fazenda.df.gov.br/aplicacoes/legislacao/legislacao/TelaSaidaDocumento.cfm?txtNumero=82&txtAno=2018&txtTipo=7&txtParte=."

	// carrega o HTML do site desejado
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// Encontre todas as linhas da tabela, exceto a primeira (cabeçalho)
	rows := doc.Find("tr").Not(":first-child")

	// Cria um arquivo CSV
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal("Erro ao criar arquivo CSV:", err)
	}
	defer file.Close()

	// Cria um escritor CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escreve o cabeçalho do CSV
	header := []string{"Nome", "CNPJ"}
	err = writer.Write(header)
	if err != nil {
		log.Fatal("Erro ao escrever cabeçalho do CSV:", err)
	}

	// Percorre cada linha da tabela e escreve as colunas no CSV
	rows.Each(func(i int, row *goquery.Selection) {
		// Encontra as colunas da linha atual
		cols := row.Find("td")

		// Obtem os valores das colunas desejadas
		nome := strings.TrimSpace(cols.Eq(0).Text())
		nome = strings.ReplaceAll(nome, "\n", "")
		cnpj := strings.TrimSpace(cols.Eq(1).Text())
		cnpj = strings.ReplaceAll(cnpj, "\n", "")

		// Escreve os valores no CSV
		err = writer.Write([]string{nome, cnpj})
		if err != nil {
			log.Fatal("Erro ao escrever linha no CSV:", err)
		}
	})

	fmt.Println("Arquivo CSV gerado com sucesso!")
}
