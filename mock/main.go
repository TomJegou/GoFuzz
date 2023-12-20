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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
