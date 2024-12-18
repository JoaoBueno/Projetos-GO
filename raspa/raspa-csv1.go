package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "http://www.fazenda.df.gov.br/aplicacoes/legislacao/legislacao/TelaSaidaDocumento.cfm?txtNumero=82&txtAno=2018&txtTipo=7&txtParte=."

	// Cria a requisição HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "text/html; charset=windows-1252")

	// Realiza a requisição HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Cria um objeto goquery a partir do corpo da resposta HTTP
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Cria um arquivo CSV para armazenar os dados
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Cria o escritor CSV e escreve o cabeçalho
	writer := csv.NewWriter(file)
	writer.Write([]string{"Nome", "CNPJ"})

	// Encontra todas as linhas da tabela, exceto a primeira (cabeçalho)
	rows := doc.Find("tr").Not(":first-child")

	// Percorre cada linha da tabela e escreve os valores das colunas desejadas no arquivo CSV
	rows.Each(func(i int, row *goquery.Selection) {
		// Encontra as colunas da linha atual
		cols := row.Find("td")

		// Extrai os valores das colunas desejadas
		nome := strings.TrimSpace(cols.Eq(0).Text())
		cnpj := strings.TrimSpace(cols.Eq(1).Text())

		// Escreve os valores no arquivo CSV
		err := writer.Write([]string{nome, cnpj})
		if err != nil {
			log.Fatal(err)
		}
	})

	// Finaliza a escrita e verifica por erros
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("CSV gravado com sucesso!")
}
