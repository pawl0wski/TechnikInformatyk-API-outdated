package main

import (
	"github.com/Jeboczek/TechnikInformatykBackend/cache"
	"github.com/Jeboczek/TechnikInformatykBackend/db"
	"github.com/Jeboczek/TechnikInformatykBackend/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize database
	backendDatabase := db.OpenDb("db.sqlite3")
	// Initialize and update cache
	cacheInstance := cache.GetCacheInstance()
	cacheInstance.UpdateCache(backendDatabase)

	// Initialize gin server
	server := gin.Default()
	// Initialize routes
	server.GET("/api/exams", routes.Exams)
	server.GET("/api/questions", routes.Questions)

	// Run server
	server.Run()

}
