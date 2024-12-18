package main

import (
	"C"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.188.1"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "fullcontrol"
)

type Xfd struct {
	// Name string		`json:"name"`
	Offset    int    `json:"offset"`
	Length    int    `json:"length"`
	Type      int    `json:"type"`
	Digits    int    `json:"digits"`
	Scale     int    `json:"scale"`
	UserType  int    `json:"usertype"`
	Condition int    `json:"condition"`
	Level     int    `json:"level"`
	Format    string `json:"format"`
}

var xfdp map[string]Xfd
var db *sql.DB

func connect() error {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	// open database
	db, err = sql.Open("postgres", psqlconn)

	if err != nil {
		return err
	}

	// check db
	err = db.Ping()
	if err != nil {
		return err
	}

	return err
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		a := []byte(s)
		switch a[len(a)-1] {
		case 125:
			a[len(a)-1] = 48
		case 74:
			a[len(a)-1] = 49
		case 75:
			a[len(a)-1] = 50
		case 76:
			a[len(a)-1] = 51
		case 77:
			a[len(a)-1] = 52
		case 78:
			a[len(a)-1] = 53
		case 79:
			a[len(a)-1] = 54
		case 80:
			a[len(a)-1] = 55
		case 81:
			a[len(a)-1] = 56
		case 82:
			a[len(a)-1] = 57
		}
		s = string(a)
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return i * -1
	}
	return i
}

/*XFDCheck is a func to check if xfd is ok*/
//export XFDCheck
func XFDCheck(arquivo *C.char, md5s *C.char, ret *int) {
	arq := C.GoString(arquivo)
	md5 := C.GoString(md5s)
	fmt.Println(arq, md5)
	if db == nil {
		err := connect()
		if err != nil {
			db = nil
			*ret = 10001 // erro de conexão
			return
		}
	}

	// close database
	defer db.Close()

	rows, err := db.Query(`SELECT "arquivo", "md5" FROM "fc_xfds"`)
	if err != nil {
		fmt.Println(err)
		*ret = 10002 // erro de leitura
		return
	}

	defer rows.Close()
	for rows.Next() {
		var arquivo string
		var md5 int
		err = rows.Scan(&arquivo, &md5)
		if err != nil {
			fmt.Println(err)
			*ret = 10003 // erro de scan
			return
		}
		fmt.Println(arquivo, md5)
	}

	*ret = 0
	return
}

/*XFDP is a func to parse a XFD*/
//export XFDP
func XFDP(str *C.char, ret *int) {
	strl := C.GoString(str)
	strs := strings.Split(strl, "\x0a")
	for i := 0; i < len(strs); i++ {
		fmt.Println(strs[i])
	}

	return
}

/*XFDParse is a func to parse a XFD*/
//export XFDParse
func XFDParse(str *C.char, ret *int) {
	strl := C.GoString(str)
	strs := strings.Split(strl, ",")
	i := len(xfdp)
	if i == 0 {
		xfdp = make(map[string]Xfd)
	}
	xfdp[strings.TrimSpace(strs[0])] = Xfd{
		// Name: strings.TrimSpace(strs[0]),
		Offset:    atoi(strs[1]),
		Length:    atoi(strs[2]),
		Type:      atoi(strs[3]),
		Digits:    atoi(strs[4]),
		Scale:     atoi(strs[5]),
		UserType:  atoi(strs[6]),
		Condition: atoi(strs[7]),
		Level:     atoi(strs[8]),
		Format:    strings.TrimSpace(strs[9])}
	*ret = i
	return
}

/*XFDtoJson is a func to convert XFD to JSON*/
//export XFDtoJson
func XFDtoJson(xfdjson **C.char, length *int) {
	// j, _ := json.MarshalIndent(xfdp, "", "    ")
	j, _ := json.Marshal(xfdp)
	var reta string
	reta = fmt.Sprint(string(j))
	// fmt.Println(reta)
	// fmt.Println(len(reta))
	*xfdjson = C.CString(reta)
	*length = len(reta)
}

/*XFDCreateTable is a func to create table from XFD*/
//export XFDCreateTable
func XFDCreateTable(xfdjson **C.char, length *int) {
	// fmt.Println(xfdp["NFC-NUMERO"])
	for k, v := range xfdp {
		fmt.Println(k, v.Offset)
	}
	fmt.Println(len(xfdp))
}

func main() {}
