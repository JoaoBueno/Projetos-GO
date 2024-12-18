package main

import (
	"encoding/json"
	"fmt"
)

type Xfd struct {
	Name string		`json:"name"`
	Offset int		`json:"offset"`
	Length int      `json:"length"`
	Type int		`json:"type"`
	Digits int		`json:"digits"`
	Scale int		`json:"scale"`
	UserType int	`json:"usertype"`
	Condition int	`json:"condition"`
	Level int		`json:"level"`
	Format string	`json:"format"`
}

const N = 10

func main() {
	datas := make(map[string]Xfd, N)
	
	for i := 0; i < 10; i++ {
		datas[fmt.Sprint(i)] = Xfd{Name: "Nome", 
		                           Offset: i,  
								   Length: i * 2, 
								   Type: i * 3,
								   Digits: i * 4,
								   Scale: i * -1,
								   UserType: i * 5,
								   Condition: i * 6,
								   Level: i * 7,
								   Format: "Format"}
	}
	// fmt.Println(datas)
	// fmt.Println()

	j, err := json.MarshalIndent(datas, "", "    ")
	fmt.Println(string(j), err)
	fmt.Println()

	datas2 := make([]Xfd, N)
	for i := 0; i < 10; i++ {
		datas2[i] = Xfd{Name: "Nome", 
						Offset: i,  
						Length: i * 2, 
						Type: i * 3,
						Digits: i * 4,
						Scale: i * -1,
						UserType: i * 5,
						Condition: i * 6,
						Level: i * 7,
						Format: "Format"}
}

// fmt.Println(datas2)
// fmt.Println()

j, err = json.MarshalIndent(datas2, "", "    ")
fmt.Println(string(j), err)
fmt.Println()



}
