package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"regexp"
	// _ "github.com/lib/pq"
)

type Field struct {
	Name string
	Type string
}

func readFD(fd string) ([]Field, error) {
	// Inicialize um slice vazio de estruturas Field
	fields := []Field{}

	// Abra o arquivo da FD para leitura
	file, err := os.Open(fd)
	if err != nil {
		return fields, err
	}
	defer file.Close()

	// Crie um scanner para ler o arquivo linha por linha
	scanner := bufio.NewScanner(file)

	// Compile uma expressão regular para extrair o nome e o tipo de um campo da FD
	// fieldRegex := regexp.MustCompile(`^\s*(\d+)\s+(\w+)\s+(\w+)\s*$`)
	fieldRegex := regexp.MustCompile(`(?:^|\n)\s*(\d+)\s+(\S+)\s+PIC\s+(\S+)(?:\s+\((\d+)(?:\s*,\s*(\d+)))`)

	// Enquanto houver linhas para serem lidas, processe as informações da FD
	for scanner.Scan() {
		// Obtenha a linha atual
		line := scanner.Text()
		// fmt.Println(line)

		// Aplique a expressão regular à linha para extrair o nome e o tipo do campo
		matches := fieldRegex.FindStringSubmatch(line)
		fmt.Println(matches)

		if len(matches) < 3 {
			continue
		}
		name := matches[2]
		fieldType := matches[3]

		// Adicione o campo à lista de campos
		fields = append(fields, Field{Name: name, Type: fieldType})
	}

	// Verifique se houve algum erro durante a leitura do arquivo
	if err := scanner.Err(); err != nil {
		return fields, err
	}

	return fields, nil
}

func createTable(db *sql.DB, tableName string, fields []Field) error {
	// Inicie a string da consulta SQL
	createTableSQL := "CREATE TABLE " + tableName + " ("

	// Adicione os campos da tabela à consulta SQL
	for i, field := range fields {
		createTableSQL += field.Name + " " + field.Type
		if i < len(fields)-1 {
			createTableSQL += ", "
		}
	}
	createTableSQL += ")"

	// Execute a consulta para criar a tabela
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// // Conecte-se ao banco de dados PostgreSQL
	// db, err := sql.Open("postgres", "user=username password=password dbname=mydatabase sslmode=disable")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer db.Close()

	// Lê a FD e cria as estruturas Field com as informações dos campos da FD
	fields, err := readFD("fd.fd1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fields)

	// // Crie a tabela
	// err = createTable(db, "eep", fields)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println("Tabela criada com sucesso!")
}
