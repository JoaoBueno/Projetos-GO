package bd

import (
	"os"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type MyError struct {
	Code    int
	Message string
}

// Error makes MyError implement the error interface.
func (e *MyError) Error() string {
    return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func Buscar(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT fco, nome, jurfis, cnpj, cidade, uf from fco ORDER BY jurfis, uf, cidade")

	if os.Getenv("DEBUG") == "true" {
		fmt.Println("SELECT fco, nome, jurfis, cnpj, cidade, uf from fco")
		// time.Sleep(3 * time.Second)
	}

	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println(err)
		}
		return nil, &MyError{Code: -1003, Message: err.Error()}
	}
	return rows, nil
}
