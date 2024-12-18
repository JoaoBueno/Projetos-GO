package main

import (
	"fmt"
	"log"
	"net"
)

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	fmt.Println(ifas)
	var as []string
	for _, ifa := range ifas {
		fmt.Println(ifa.Name)
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}

func main() {
	as, err := getMacAddr()
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range as {
		fmt.Println(a)
	}
}
