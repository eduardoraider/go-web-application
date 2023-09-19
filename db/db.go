package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DatabaseConnection() *sql.DB {
	conn := "user=postgres dbname=golang_store password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
