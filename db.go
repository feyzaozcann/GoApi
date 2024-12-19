package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDatabase() error {
	var err error
	DB, err = sql.Open("postgres", "host=localhost user=postgres password=12345 dbname=showsdb sslmode=disable")

	if err != nil {
		return err
	}
	return nil
}

func CloseDatabase() error {
	return DB.Close()
}
