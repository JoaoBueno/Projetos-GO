package main

import "fmt"

func main() {
	a := 10
	b := "World"
	c := 3.1415
	d := false
	e := `isisisii
	iisisiis
	isisiis`

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)
}
