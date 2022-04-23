package db

import (
	"database/sql"
	"log"
)

func GetImage(backendDatabase *sql.DB, questionUuid string) []byte {
	row, err := backendDatabase.Query("SELECT image FROM question WHERE uuid = ?", questionUuid)
	defer row.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var image []byte
	row.Next()
	err = row.Scan(&image)
	if err != nil {
		return []byte{}
	}
	return image
}
