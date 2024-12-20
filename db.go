package main

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func OpenDatabase() error {
	var err error
	DB, err = sqlx.Open("postgres", "host=localhost user=postgres password=12345 dbname=showsdb sslmode=disable")

	if err != nil {
		return err
	}
	return nil
}

func CloseDatabase() error {
	return DB.Close()
}
