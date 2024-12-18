package main

import "fmt"

type Carro struct {
	Name string
}

func (f Carro) andou() {
	f.Name = "BMW"
	fmt.Println(f.Name, "andou")
}

func (p *Carro) pandou() {
	p.Name = "BMW"
	fmt.Println(p.Name, "andou")
}

func main() {
	carro := Carro{
		Name: "Ka",
	}

	carro.andou()
	fmt.Println(carro.Name)

	carro.pandou()
	fmt.Println(carro.Name)

}
