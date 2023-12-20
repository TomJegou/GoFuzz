package main

import (
	"fmt"
	"log"
	"net/http"
)

type Endpoints struct {
	Route    string
	Activate bool
}

const (
	port = "8080"
)

var (
	listRoutes = []Endpoints{
		{
			Route:    "/",
			Activate: true,
		},
		{
			Route:    "/robot.txt",
			Activate: true,
		},
	}
)

func main() {
	for _, route := range listRoutes {
		http.HandleFunc(route.Route, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Route path ok for : %s", route.Route)
		})
	}
	log.Printf("Mock server is up !\nListening on : http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
