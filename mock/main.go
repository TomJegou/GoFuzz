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
			Route:    "/home",
			Activate: true,
		},
		{
			Route:    "/robot",
			Activate: true,
		},
	}
)

func main() {
	for i := range listRoutes {
		route := listRoutes[i]
		if route.Activate {
			http.HandleFunc(route.Route, func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Route path ok for : %s", route.Route)
			})
		}
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	log.Printf("Mock server is up ! 🔥\nListening on : http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
