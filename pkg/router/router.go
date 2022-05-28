package router

import (
	"log"
	"net/http"
	"time"

	"github.com/Becram/sql-prometheus-metrics/pkg/middleware"
	"github.com/Becram/sql-prometheus-metrics/pkg/models"
	"github.com/gorilla/mux"
)

type Routes []models.Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/getallrunning", middleware.GetAllRunningJobs).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getallfailed", middleware.GetAllFailedJobs).Methods("GET", "OPTIONS")

	return router
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
