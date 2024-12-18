package main

import (
	"calc/calc"
	"fmt"
)

func main() {
	var n1 = 4
	n2 := 0
	r, err := calc.Dividir(n1, n2)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println(r)
}
