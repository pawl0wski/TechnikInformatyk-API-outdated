package db

import (
	"database/sql"
	"fmt"
	"log"
)

func OpenDbWithDefaultConnectionPath() *sql.DB {
	return OpenDb(GetDbConnectionPath())
}

func OpenDb(connectionPath string) *sql.DB {
	fmt.Println(connectionPath)
	db, err := sql.Open("mysql", connectionPath)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
