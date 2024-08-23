package dbdemo

import (
	"database/sql"
	"time"
)

func InsertDemoRecord(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO demo_record (created_at) VALUES ($1)", time.Now())
	return err
}

func CountDemoRecords(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM demo_record").Scan(&count)
	return count, err
}
