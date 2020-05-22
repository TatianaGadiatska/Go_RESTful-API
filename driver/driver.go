package driver

import (
	"database/sql"
	"log"
)

var db *sql.DB
var connStr string = "user=postgres password=123 dbname=books-store sslmode=disable"

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", connStr)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
