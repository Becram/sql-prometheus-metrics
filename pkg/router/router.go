package event_observers

import (
	"log"
	"net/http"
	"time"

	"github.com/Becram/sql-prometheus-metrics/pkg/event_observers"
	"github.com/Becram/sql-prometheus-metrics/pkg/models"
	"github.com/gorilla/mux"
)

type Routes []models.Route

var routes = Routes{
	models.Route{
		"jobs",
		"GET",
		"/",
		even,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

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
