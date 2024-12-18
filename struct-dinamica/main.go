package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

const N = 10

func main() {

	m := make(map[string]int)
	for i := 0; i < 10; i++ {
		m[fmt.Sprint(i)] = i * i
	}

	fmt.Println(m)

	j, err := json.Marshal(m)
	fmt.Println(string(j), err)
	fmt.Println()
	
	fmt.Println(m["9"])

	//option 1
	datas := make(map[string]Foo, N)

	for i := 0; i < 10; i++ {
		datas[fmt.Sprint(i)] = Foo{Number: i, Title: "test"}
	}
	fmt.Println(datas)
	fmt.Println()

	j, err = json.Marshal(datas)
	fmt.Println(string(j), err)
	fmt.Println()

	//option 2
	datas2 := make([]Foo, N)
	for i := 0; i < 10; i++ {
		datas2[i] = Foo{Number: i, Title: "test"}
	}
	fmt.Println(datas2)
	fmt.Println()
	j, err = json.Marshal(datas2)
	fmt.Println(string(j), err)
}

