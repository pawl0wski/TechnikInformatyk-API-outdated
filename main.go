package main

import (
	"fmt"
	"github.com/pawl0wski/technikinformatyk-backend/cdn"
	"github.com/pawl0wski/technikinformatyk-backend/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/pawl0wski/technikinformatyk-backend/db"
	dblocker "github.com/pawl0wski/technikinformatyk-backend/db_locker"
)

func loadDotFile() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
}

func initializeDBLocker() {
	database := db.OpenDbWithDefaultConnectionPath()
	dbLocker := dblocker.GetDBLockerInstance()
	dbLocker.DB = database
}

func main() {
	loadDotFile()

	// Initialize gin server
	server := gin.Default()

	initializeDBLocker()
	cdn.RebuildCdn()
	routes.SetupRoutes(server)

	// Run server
	server.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))

}
