package bd

import (
	"database/sql"
	"fmt"
	"strconv"
)

type Municipio struct {
	CodigoIbge   int
	Nome         string
	UF           string
	CodigoSerpro int
}

func BuscarMunicipio(db *sql.DB, codigo_ibge int) (Municipio, error) {
	_, create_err := db.Query(`CREATE TABLE IF NOT EXISTS municipios (
															codigo_ibge INTEGER PRIMARY KEY,
															nome VARCHAR NOT NULL,
															uf VARCHAR(2) NOT NULL,
															codigo_serpro INTEGER NOT NULL
														)`)

	if create_err != nil {
		fmt.Println(create_err)
		return Municipio{}, create_err
	}

	municipioBd, err := db.Query(`SELECT * FROM "municipios" WHERE "codigo_ibge" = '` + strconv.Itoa(codigo_ibge) + `'`)

	if err != nil {
		fmt.Println(err)
		return Municipio{}, err
	}

	defer municipioBd.Close()

	var municipio Municipio
	if municipioBd.Next() {
		var codigo_ibge int
		var nome string
		var uf string
		var codigo_serpro int

		err = municipioBd.Scan(&codigo_ibge, &nome, &uf, &codigo_serpro)

		if err != nil {
			fmt.Println(err)
			return Municipio{}, err
		}

		municipio = Municipio{
			CodigoIbge:   codigo_ibge,
			Nome:         nome,
			UF:           uf,
			CodigoSerpro: codigo_serpro,
		}
	}

	if err != nil {
		fmt.Println(err)
		return Municipio{}, err
	}

	return municipio, nil

}
