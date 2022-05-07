package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Becram/sql-prometheus-metrics/pkg/models"
	"github.com/Becram/sql-prometheus-metrics/pkg/router"
)

type Routes []models.Route

var routes = Routes{
	models.Route{
		"jobs",
		"GET",
		"/",
		event_observers.jobHandler,
	},
}

func main() {

	router := router.NewRouter()
	fmt.Print("Serving http request at localhost:8080.....\n")
	log.Fatal(http.ListenAndServe(":8080", router))

}
