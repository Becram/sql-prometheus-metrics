package event_observers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Becram/sql-prometheus-metrics/pkg/models"
	_ "github.com/lib/pq"
)

func appInit() {

	if os.Getenv("DB_HOST") == "" {
		fmt.Fprintln(os.Stderr, "Please provide 'DB_HOST' through environment")
		os.Exit(1)
	}

	if os.Getenv("DB_PORT") == "" {
		fmt.Fprintln(os.Stderr, "Please provide 'DB_PORT' through environment")
		os.Exit(1)
	}
	if os.Getenv("DB_USER") == "" {
		fmt.Fprintln(os.Stderr, "Please provide 'DB_USER' through environment")
		os.Exit(1)
	}
	if os.Getenv("DB_PASSWORD") == "" {
		fmt.Fprintln(os.Stderr, "Please provide 'DB_PASSWORD' through environment")
		os.Exit(1)
	}
	if os.Getenv("DB_NAME") == "" {
		fmt.Fprintln(os.Stderr, "Please provide 'DB_NAME' through environment")
		os.Exit(1)
	}

}

var (
	hostname      = os.Getenv("DB_HOST")
	host_port     = os.Getenv("DB_PORT")
	username      = os.Getenv("DB_USER")
	password      = os.Getenv("DB_PASSWORD")
	database_name = os.Getenv("DB_NAME")
)

func jobHandler(w http.ResponseWriter, r *http.Request) {
	appInit()
	// fmt.Printf("Dial to %s:%d\n", hostname, host_port)
	// err := telnet.DialToAndCall(fmt.Sprintf("%s:%d", hostname, host_port), caller{})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	pg_con_string := fmt.Sprintf("port=%s host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, database_name)

	// fmt.Printf("%s\n", pg_con_string)
	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}

	fmt.Println("Successfully connected!")

	rows, err := db.Query("SELECT * FROM public.event_observers where status = 'running' ")
	if err != nil {
		log.Fatalf("could not execute query: %v", err)
	}

	events := []models.EventObservers{}

	for rows.Next() {
		event := models.EventObservers{}
		if err := rows.Scan(&event.ID, &event.Status, &event.ForceStop, &event.CreatedAt, &event.UpdatedAt, &event.JobType); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		events = append(events, event)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(events); err != nil {
		panic(fmt.Errorf("failed to get status: %v", err))
	}

	fmt.Printf("found %d jobs running\n", len(events))

}
