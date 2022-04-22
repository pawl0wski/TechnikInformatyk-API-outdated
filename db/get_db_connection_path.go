package db

import "os"

func GetDbConnectionPath() string {
	connectionPath := os.Getenv("MYSQL_CONNECTION_PATH")
	return connectionPath
}
