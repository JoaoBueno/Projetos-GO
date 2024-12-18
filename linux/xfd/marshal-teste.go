package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	data := map[string]int{
		"b": 2,
		"a": 1,
	}

	json, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}
