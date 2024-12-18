package main

import "fmt"

func main() {
	a := 10
	var ponteiro *int = &a
	*ponteiro = 50
	b := &a
	*b = 60
	fmt.Println(a)

	variavel := 10
	abc(&variavel)
	fmt.Println(variavel)
}

func abc(a *int) {
	*a = 200
}
