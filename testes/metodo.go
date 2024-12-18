package main

import "fmt"

type Carro struct {
	Nome string
}

func (c Carro) andar() {
	fmt.Println(c.Nome, "andou!")
}

func main() {
	carro := Carro{Nome: "Bmw"}
	carro.andar()
}
