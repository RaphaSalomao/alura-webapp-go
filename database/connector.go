package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const datasource = "host=localhost port=5433 user=postgres dbname=postgres password=root sslmode=disable"

func Connect() *sql.DB {
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		panic(err)
	}
	return db
}
