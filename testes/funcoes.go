package main

import "fmt"

func main() {
	res := somaTudo(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(res)
}

func somaTudo(x ...int) int {
	result := 0
	for _, v := range x {
		result += v
	}
	return result
}
