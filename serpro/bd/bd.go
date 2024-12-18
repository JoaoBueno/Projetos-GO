package bd

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "FullControl"
)

func Connect() (*sql.DB, error) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// var err error
	// open database
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		return nil, err
	}

	// check db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

func Verifica(db *sql.DB, cnpj string) error {
	rows, err := db.Query(`SELECT "cnpj", "data", "md5" FROM "consultas_serpro" WHERE "cnpj" = '` + cnpj + `'`)
	fmt.Println(rows, err)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var cnpj string
		var data time.Time
		var md5 string
		err = rows.Scan(&cnpj, &data, &md5)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(cnpj, "-", data, "-", md5)
	}

	return err

}
