package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jpawlowskii/TechnikInformatykBackend/cache"
	"github.com/jpawlowskii/TechnikInformatykBackend/db"
	"github.com/jpawlowskii/TechnikInformatykBackend/routes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize database
	backendDatabase := db.OpenDb("db.sqlite3")
	// Initialize and update cache
	cacheInstance := cache.GetCacheInstance()
	cacheInstance.UpdateCache(backendDatabase)

	// Initialize gin server
	server := gin.Default()
	// Enable gzip
	server.Use(gzip.Gzip(gzip.DefaultCompression))
	// Initialize routes
	server.GET("/api/exams", routes.Exams)
	server.GET("/api/questions", routes.Questions)
	server.GET("/api/image/:questionUuid", routes.Image)
	server.GET("/api/databaseVersion", routes.DatabaseVersion)

	// Run server
	server.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))

}
