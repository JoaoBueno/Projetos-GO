package main

import (
	"database/sql"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"serpro/bd"
	"serpro/md5"
	"strings"
)

var db *sql.DB

// Consumer Key : rs_f7yOqYa9K5s45kkun6EBqwcMa
// Consumer Secret : xhVqUCZ69Pyjwo6aIp2PDWd8py8a

func main() {
	fmt.Println("Starting the application...")

	teste := "bunda"

	teste = os.Args[1]
	fmt.Println(teste)

	url := "https://gateway.apiserpro.serpro.gov.br/token"
	authKey := "rs_f7yOqYa9K5s45kkun6EBqwcMa" + ":" + "xhVqUCZ69Pyjwo6aIp2PDWd8py8a"
	sEnc := b64.StdEncoding.EncodeToString([]byte(authKey))
	postString := "grant_type=client_credentials"

	// fmt.Println("URL:>", url)
	// fmt.Println(authKey)
	// fmt.Println(sEnc)

	cliente := &http.Client{}
	req, _ := http.NewRequest("POST", url, strings.NewReader(postString))
	req.Header.Add("Authorization", "Basic "+sEnc)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := cliente.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	fmt.Println("response Headers:", res.Header)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response Body:", string(body))

	var result map[string]interface{}
	var at = ""
	json.Unmarshal([]byte(body), &result)

	for key, value := range result {
		if key == "access_token" {
			at = fmt.Sprintf("%v", value)
		}
		fmt.Println(key, value)
	}

	fmt.Println(at)

	// req, _ = http.NewRequest("GET", "https://gateway.apiserpro.serpro.gov.br/consulta-cnpj-df/v2/basica/34238864000168", nil)
	req, _ = http.NewRequest("GET", "https://gateway.apiserpro.serpro.gov.br/consulta-cnpj-df/v2/basica/24907602000195", nil)
	// req.Header.Add("Authorization", "Bearer 06aef429-a981-3ec5-a1f8-71d38d86481e")
	req.Header.Add("Authorization", "Bearer "+at)
	req.Header.Add("accept", "application/json")

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println(req)
	fmt.Println("-----------------------------------------------------------------------------------")

	res, err = cliente.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	fmt.Println("response Headers:", res.Header)
	body, _ = ioutil.ReadAll(res.Body)
	fmt.Println("response Body:", string(body))

	fmt.Println("Terminating the application...")

	// a := "{\"nome\":\"Bueno\",\"idade\":56}"
	a := string(body)

	reta := ""
	length := 0
	md5.MD5String(a, &reta, &length)

	fmt.Println(a, reta, length)

	// if 1==1 {return}

	// curl -X GET "https://gateway.apiserpro.serpro.gov.br/consulta-cnpj-trial/v1/cnpj/34238864000168" \
	// -H "accept: application/json" \
	// -H "Authorization: Bearer 06aef429-a981-3ec5-a1f8-71d38d86481e" \
	// curl -X GET "https://gateway.apiserpro.serpro.gov.br/consulta-cnpj-trial/v1/cnpj/34238864000168" -H "accept: application/json" -H "Authorization: Bearer e0e79ecf-3c31-3c17-88c7-4ff2e3d688b7"

	// curl -X GET "https://gateway.apiserpro.serpro.gov.br/consulta-cnpj-trial/v1/cnpj/34238864000168" \
	// -H "accept: application/json" \
	// -H "Authorization: Bearer 06aef429-a981-3ec5-a1f8-71d38d86481e" \

	if db == nil {
		db, err = bd.Connect()
		if err != nil {
			db = nil
			fmt.Println("Erro ao conectar no banco de dados, " + err.Error())
			return
		}
	}

	// close database
	defer db.Close()

	fmt.Println("Conectado")

	err = bd.Verifica(db, "24907602000195")

}
