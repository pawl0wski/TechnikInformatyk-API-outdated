package db

import (
	"os"

	"github.com/joho/godotenv"
)

func GetDbConnectionPath() string {
	connectionPath := os.Getenv("MYSQL_CONNECTION_PATH")
	if len(connectionPath) == 0 {
		godotenv.Load("../.env")
		connectionPath = os.Getenv("MYSQL_CONNECTION_PATH")
	}
	return connectionPath
}
