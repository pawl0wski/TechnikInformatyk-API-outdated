package db

import (
	"database/sql"
	"log"
)

func OpenDb(fileName string) *sql.DB {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
