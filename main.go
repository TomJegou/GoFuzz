package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	listArgs := os.Args[1:]
	// Defensive checking for the host
	if len(listArgs) == 0 {
		fmt.Println("🚫 No host provided")
		return
	}
	host := listArgs[0]
	// Defensive checking for the keyword fuzz
	if !strings.Contains(host, "FUZZ") {
		fmt.Println("🚫 No keyword FUZZ found")
		return
	}
	// Parse the host parameter
	listEndpointParts := strings.Split(host, "FUZZ")
	byteDic, err := os.ReadFile("./dic.txt")
	if err != nil {
		panic(fmt.Errorf("os.ReadFile('./dictionnary.txt') Got : %s", err.Error()))
	}
	// Parse the dictionnary
	words := strings.Split(string(byteDic), "\n")
	for _, word := range words {
		target := ""
		for i := 0; i < len(listEndpointParts)-1; i++ {
			// Create the url with the endpoit from de dictionnary
			target += listEndpointParts[i] + word
		}
		// Create the request
		req, err := http.NewRequest(http.MethodGet, target, nil)
		if err != nil {
			panic(err)
		}
		// Send the request
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		// Check the status code
		if res.StatusCode == 200 {
			fmt.Printf("Endpoint found : %s\n", target)
		}
	}

}
