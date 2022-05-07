package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/Becram/k8s-api-client/pkg/k8s"
	"github.com/Becram/sql-prometheus-metrics/pkg/router"
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
		events_,
	},
	Route{
		"restartDeployment",
		"POST",
		"/restart",
		k8s.RestartDeployment,
	},
	Route{
		"listDeployment",
		"POST",
		"/list",
		k8s.ListDeployment,
	},
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		event_observers.jobHandler,
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
