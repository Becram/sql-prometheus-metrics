package middleware

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

// create connection with postgres db
func createConnection() *sql.DB {
	appInit()

	pg_con_string := fmt.Sprintf("port=%s host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, database_name)

	// fmt.Printf("%s\n", pg_con_string)
	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

func GetAllRunningJobs(w http.ResponseWriter, r *http.Request) {

	// get all the users in the db
	events, err := getAllRunningJobs()

	if err != nil {
		log.Fatalf("Unable to get all running jobs. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(events)

}

//------------------------- handler functions ----------------
func getAllRunningJobs() ([]models.Event, error) {
	db := createConnection()

	// close the db connection
	defer db.Close()

	var events []models.Event

	// create the select sql query
	sqlStatement := `SELECT * FROM public.event_observers where status = 'running'`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	for rows.Next() {
		var event models.Event

		// unmarshal the row object to user
		err = rows.Scan(&event.ID, &event.Status, &event.ForceStop, &event.CreatedAt, &event.UpdatedAt, &event.JobType)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		events = append(events, event)
	}

	return events, err

}
