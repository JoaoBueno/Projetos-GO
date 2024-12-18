package bd

import (
	"os"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type RegistroSerpro struct {
	Cnpj  string
	Data  time.Time
	Md5   string
	Dados string
}

type MyError struct {
	Code    int
	Message string
}

// Error makes MyError implement the error interface.
func (e *MyError) Error() string {
    return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func Buscar(db *sql.DB, cnpj string) (RegistroSerpro, error) {
	_, create_err := db.Query(`CREATE TABLE IF NOT EXISTS consultas_serpro (
															cnpj VARCHAR(14) PRIMARY KEY, 
															data DATE, 
															md5 VARCHAR(32), 
															dados JSONB)`)

	if create_err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println(create_err.Error())
		}
		return RegistroSerpro{}, &MyError{Code: -1002, Message: create_err.Error()}
	}

	registro, err := db.Query(`SELECT "cnpj", "data", "md5", "dados" FROM "consultas_serpro" WHERE "cnpj" = '` + cnpj + `'`)

	if os.Getenv("DEBUG") == "true" {
		fmt.Println(cnpj, err)
		fmt.Println(`SELECT "cnpj", "data", "md5", "dados" FROM "consultas_serpro" WHERE "cnpj" = '` + cnpj + `'`)
		// time.Sleep(3 * time.Second)
	}

	if err != nil {
		fmt.Println(err)
		return RegistroSerpro{}, &MyError{Code: -1003, Message: create_err.Error()}
	}

	defer registro.Close()

	var registroSerpro RegistroSerpro
	if registro.Next() {
		var cnpj1 string
		var data time.Time
		var md5 string
		var dados string
		err = registro.Scan(&cnpj1, &data, &md5, &dados)

		if os.Getenv("DEBUG") == "true" {
			fmt.Println(cnpj, cnpj1, err)
			// time.Sleep(3 * time.Second)
		}
	
		if err != nil {
			if os.Getenv("DEBUG") == "true" {
				fmt.Println(err.Error())
			}
			return RegistroSerpro{}, &MyError{Code: -1004, Message: err.Error()}
		}

		registroSerpro = RegistroSerpro{Cnpj: cnpj1, Data: data, Md5: md5, Dados: dados}
	}

	return registroSerpro, nil
}

func Inserir(db *sql.DB, dado RegistroSerpro) error {
	stmt, err := db.Prepare("INSERT INTO consultas_serpro(cnpj, data, md5, dados) VALUES($1, $2, $3, $4)")
	
	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println(err.Error())
		}
		return &MyError{Code: -1005, Message: err.Error()}
	}
	defer stmt.Close()

	_, err = stmt.Exec(dado.Cnpj, dado.Data, dado.Md5, dado.Dados)

	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println(err.Error())
		}
		return &MyError{Code: -1006, Message: err.Error()}
	}

	return nil
}

func Atualizar(db *sql.DB, dado RegistroSerpro) error {
	stmt, err := db.Prepare("UPDATE consultas_serpro SET data = $1, md5 = $2, dados = $3 WHERE cnpj = $4")

	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println(err.Error())
		}
		return &MyError{Code: -1007, Message: err.Error()}
	}
	defer stmt.Close()

	_, err = stmt.Exec(dado.Data, dado.Md5, dado.Dados, dado.Cnpj)

	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println(err.Error())
		}
		return &MyError{Code: -1008, Message: err.Error()}
	}

	return nil
}

func AtualizarData(db *sql.DB, dado RegistroSerpro) error {
	stmt, err := db.Prepare("UPDATE consultas_serpro SET data = $1 WHERE cnpj = $2")

	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println(err.Error())
		}
		return &MyError{Code: -1009, Message: err.Error()}
	}
	defer stmt.Close()

	_, err = stmt.Exec(dado.Data, dado.Cnpj)

	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println(err.Error())
		}
		return &MyError{Code: -1010, Message: err.Error()}
	}

	return nil
}
