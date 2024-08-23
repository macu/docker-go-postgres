package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"docker-go-postgres/pkg/dbdemo"
)

func main() {

	// Connect to postgres, using DB_HOST, DB_USER, DB_PASS, DB_NAME
	dbUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// Insert new demo_record
		err := dbdemo.InsertDemoRecord(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s\n", err)
		}

		// Count entries in demo_record
		count, err := dbdemo.CountDemoRecords(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s\n", err)
		}

		// Print count
		fmt.Fprintf(w, "Count: %d\n", count)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write link to /test
		fmt.Fprintf(w, "<a href='/test'>Test</a>")
	})

	http.ListenAndServe(":8080", nil)

}
