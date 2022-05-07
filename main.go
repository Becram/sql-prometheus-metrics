package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/Becram/sql-prometheus-metrics/pkg/event_observers"
	"github.com/Becram/sql-prometheus-metrics/pkg/home"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route


var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		events.,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {

	router := NewRouter()
	fmt.Print("Serving http request at localhost:8080.....\n")
	log.Fatal(http.ListenAndServe(":8080", router))

}
