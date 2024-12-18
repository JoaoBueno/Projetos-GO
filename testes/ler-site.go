package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://google.com.br")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res.Header)

	res, err = http.Get("https://sitemuitodoido.com.br")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res.Header)

}
