package main

import (
	"github.com/Jeboczek/TechnikInformatykBackend/cache"
	"github.com/Jeboczek/TechnikInformatykBackend/db"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize database
	backendDatabase := db.OpenDb("db.sqlite3")
	// Initialize and update cache
	cache.GetCacheInstance().UpdateCache(backendDatabase)
}
