package bd

import (
	"database/sql"
	"fmt"
	"os"
)

type Cnae struct {
	Id        string
	Descricao string
}

func BuscarCnae(db *sql.DB, id string) (Cnae, error) {
	cnaeBd, err := db.Query(`SELECT * FROM "cnae" WHERE "id" = $1`, id)
	
	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println(err.Error())
		}
		return Cnae{}, &MyError{Code: -1011, Message: err.Error()}
	}

	defer cnaeBd.Close()

	if err != nil {
		fmt.Println("Erro na busca do CNAE: ", err)
		return Cnae{}, err
	}

	var cnae Cnae
	if cnaeBd.Next() {
		var id string
		var descricao string

		err = cnaeBd.Scan(&id, &descricao)

		if err != nil {
			fmt.Println(err)
			return Cnae{}, err
		}

		cnae = Cnae{
			Id:        id,
			Descricao: descricao,
		}
	}

	if err != nil {
		fmt.Println(err)
		return Cnae{}, err
	}

	return cnae, nil
}

func InserirCnae(db *sql.DB, cnae Cnae) error {
	cnaeBd, err := db.Query(`SELECT * FROM "cnae" WHERE "id" = $1`, cnae.Id)
	
	if err != nil {
		fmt.Println("Erro na busca do CNAE: ", err)
		return err
	}

	defer cnaeBd.Close()

	var cnaeBuscado Cnae
	if cnaeBd.Next() {
		var id string
		var descricao string

		err = cnaeBd.Scan(&id, &descricao)

		if err != nil {
			fmt.Println(err)
			return err
		}

		cnaeBuscado = Cnae{
			Id:        id,
			Descricao: descricao,
		}
	}

	if cnaeBuscado.Id == "" || cnaeBuscado.Descricao == "" {
		fmt.Println("CNAE não encontrado na base de dados, inserindo...")

		cnaeBd, err := db.Query("INSERT INTO cnae(id, descricao) VALUES($1, $2)", cnae.Id, cnae.Descricao)
		
		if err != nil {
			fmt.Println("Erro na inserção da CNAE: ", err.Error())
			return err
		}

		defer cnaeBd.Close()

	}

	return nil
}
