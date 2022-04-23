package db

import (
	"database/sql"
	"log"

	"github.com/jpawlowskii/TechnikInformatykBackend/structs"
)

func calculateChecksumOfTable(backendDatabase *sql.DB, tableName string) uint32 {
	row, err := backendDatabase.Query("CHECKSUM TABLE " + tableName)
	defer row.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var table string
	var tableChecksum []uint8
	row.Next()
	err = row.Scan(&table, &tableChecksum)
	if err != nil {
		log.Fatalln(err)
	}

	var finalResult uint32 = 0
	for _, v := range tableChecksum {
		if finalResult == 0 {
			finalResult = uint32(v)
		} else {
			finalResult *= uint32(v)
		}
	}
	return finalResult
}

func CalculateDatabaseVersion(backendDatabase *sql.DB) structs.DatabaseVersion {

	var databaseChecksum uint32
	tablesToCalculate := [3]string{"question", "exam", "questionToExam"}
	for _, v := range tablesToCalculate {
		databaseChecksum += calculateChecksumOfTable(backendDatabase, v)
	}

	return structs.DatabaseVersion{Version: databaseChecksum}
}
