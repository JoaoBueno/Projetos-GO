package apiSerpro

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type MyError struct {
	Code    int
	Message string
}

// Error makes MyError implement the error interface.
func (e *MyError) Error() string {
    return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// BuscarCnpjSerpro is a function that searches for a CNPJ in the SERPRO API and returns the company's data.
//
// Parameters:
// - cnpj (string): The CNPJ to search for.
//
// Returns:
// - string: The company's data as a string.
// - error: An error if there was an issue with the API request.
func BuscarCnpjSerpro(cnpj string) (string, error) {
	url := "https://gateway.apiserpro.serpro.gov.br/token"
	authKey := os.Getenv("CONSUMER_KEY") + ":" + os.Getenv("CONSUMER_SECRET")
	sEnc := b64.StdEncoding.EncodeToString([]byte(authKey))
	postString := "grant_type=client_credentials"

	cliente := &http.Client{}
	req, _ := http.NewRequest("POST", url, strings.NewReader(postString))
	req.Header.Add("Authorization", "Basic "+sEnc)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := cliente.Do(req)
	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println("Erro ao buscar token: " + err.Error())
		}
		return "", &MyError{Code: -2001, Message: err.Error()}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println("Erro ao ler o token: " + err.Error())
		}
		return "", &MyError{Code: -2002, Message: err.Error()}
	}

	var result map[string]interface{}
	var token = ""
	json.Unmarshal([]byte(body), &result)

	for key, value := range result {
		if key == "access_token" {
			token = fmt.Sprintf("%v", value)
		}
	}

	req, _ = http.NewRequest("GET", "https://gateway.apiserpro.serpro.gov.br/consulta-cnpj-df/v2/basica/"+cnpj, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("accept", "application/json")

	res, err = cliente.Do(req)

	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println("Erro ao realizar a requisição na API SERPRO: " + err.Error())
		}
		return "", &MyError{Code: -2003, Message: err.Error()}
	}

	if res.StatusCode != 200 {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println("Erro ao buscar dados da API SERPRO Status " + res.Status)
		}
		return "", &MyError{Code: -2004, Message: res.Status + " - " + strconv.Itoa(res.StatusCode)}
	}

	body, err = io.ReadAll(res.Body)

	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			fmt.Println("Erro ao ler o token: " + err.Error())
		}
		return "", &MyError{Code: -2005, Message: err.Error()}
	}

	dadosEmpresa := string(body)

	return dadosEmpresa, err
}
