package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jpawlowskii/TechnikInformatykBackend/cache"
	"github.com/jpawlowskii/TechnikInformatykBackend/cdn"
	"github.com/jpawlowskii/TechnikInformatykBackend/db"
	"github.com/jpawlowskii/TechnikInformatykBackend/routes"

	_ "github.com/mattn/go-sqlite3"
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
	backendDatabase := db.OpenDb("db.sqlite3")
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
