package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Becram/sql-prometheus-metrics/pkg/router"
)

func main() {

	r := router.NewRouter()
	fmt.Print("Serving http request at localhost:8080.....\n")
	log.Fatal(http.ListenAndServe(":8080", r))

}
