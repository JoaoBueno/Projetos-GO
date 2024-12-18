package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)
 
const (
    host     = "192.168.188.1"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "fullcontrol"
)

var db *sql.DB
 
func main() {
        // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    
	var err error

	fmt.Println(db)

        // open database
    db, err = sql.Open("postgres", psqlconn)
	fmt.Printf("%v, %T\n", psqlconn, psqlconn)
	fmt.Printf("%v, %T\n", db, db)
	fmt.Printf("%v, %T\n", err, err)
    CheckError(err)
     
        // close database
    defer db.Close()
 
        // check db
    err = db.Ping()
    CheckError(err)
 
    fmt.Println("Connected!")


    // // insert
    // // hardcoded
    // insertStmt := `insert into "fc_xfds"("arquivo", "md5", "xfd-parse") values('aivenfcp', '121', 'teste')`
    // _, e := db.Exec(insertStmt)
    // CheckError(e)
 
    // // dynamic
    // insertDynStmt := `insert into "fc_xfds"("arquivo", "md5", "xfd-parse") values($1, $2, $3)`
    // _, e = db.Exec(insertDynStmt, "aivenfdp", "1111", "tetetetetetetete")
    // CheckError(e)

	rows, err := db.Query(`SELECT * FROM "fc_xfds"`)
    CheckError(err)
 
	defer rows.Close()
	for rows.Next() {
		var arquivo string
		var md5 string
		var xfd_parse string
 
	    err = rows.Scan(&arquivo, &md5, &xfd_parse)
    	CheckError(err)
 
    	fmt.Println(arquivo, md5, xfd_parse)
	}
 
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}