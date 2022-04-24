package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pawl0wski/technikinformatyk-backend/cache"
	"github.com/pawl0wski/technikinformatyk-backend/cdn"
	"github.com/pawl0wski/technikinformatyk-backend/db"
	"github.com/pawl0wski/technikinformatyk-backend/routes"

	_ "github.com/go-sql-driver/mysql"
)

func loadDotFile() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
}

func setupRoutes(server *gin.Engine) {
	server.GET("/api/exams", routes.Exams)
	server.GET("/api/questions", routes.Questions)
	server.GET("/api/image/:questionUuid", routes.Image)
	server.GET("/api/databaseVersion", routes.DatabaseVersion)
	server.GET("/api/ping", routes.Ping)
	server.NoRoute(routes.NotFound)
}

func setupCache() {
	// Initialize database
	backendDatabase := db.OpenDbWithDefaultConnectionPath()
	// Initialize and update cache
	cacheInstance := cache.GetCacheInstance()
	cacheInstance.UpdateCache(backendDatabase)
}

func main() {
	loadDotFile()

	// Initialize gin server
	server := gin.Default()

	setupCache()
	cdn.RebuildCdn()
	setupRoutes(server)

	// Run server
	server.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))

}
