package db

import (
	"os"

	"github.com/joho/godotenv"
)

func GetDbConnectionPath() string {
	if _, ok := os.LookupEnv("MYSQL_CONNECTION_PATH"); !ok {
		godotenv.Load("../.env")
	}
	connectionPath := os.Getenv("MYSQL_CONNECTION_PATH")
	return connectionPath
}
