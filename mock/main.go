package main

import (
	"fmt"
	"log"
	"net/http"
)

// Endpoint structure
type Endpoints struct {
	Route    string
	Activate bool
}

const (
	port = "8080"
)

var (
	// List of routes to activate or deactivate
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
	// Handle the 404 error
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Get request for : http://0.0.0.0:%s%s", port, r.URL)
		http.NotFound(w, r)
	})
	log.Printf("Mock server is up ! 🔥\nListening on : http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
