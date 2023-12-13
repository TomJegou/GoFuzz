package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) == 0 {
		panic("No endpoint provided !")
	}
	listArgs := os.Args[1:]
	host := listArgs[0]
	listEndpointParts := strings.Split(host, "FUZZ")
	byteDic, err := os.ReadFile("./dic.txt")
	if err != nil {
		panic(fmt.Errorf("os.ReadFile('./dic.txt') Got : %s", err.Error()))
	}
	words := strings.Split(string(byteDic), "\n")
	for _, word := range words {
		target := ""
		for i := 0; i < len(listEndpointParts)-1; i++ {
			target += listEndpointParts[i] + word
		}
		fmt.Println(target)
	}
}
