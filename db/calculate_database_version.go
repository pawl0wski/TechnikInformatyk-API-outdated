package db

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jpawlowskii/TechnikInformatykBackend/structs"
)

func CalculateDatabaseVersion() structs.DatabaseVersion {
	file, err := os.Open("db.sqlite3")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		log.Fatalln(err)
	}

	return structs.DatabaseVersion{Version: fmt.Sprintf("%x", hash.Sum(nil))}
}
