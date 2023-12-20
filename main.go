package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	listArgs := os.Args[1:]
	if len(listArgs) == 0 {
		fmt.Println("⚠️No host provided")
		return
	}
	host := listArgs[0]
	if !strings.Contains(host, "FUZZ") {
		fmt.Println("⚠️No keyword FUZZ found")
		return
	}
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

	}

}
