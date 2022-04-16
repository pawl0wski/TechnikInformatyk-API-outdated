package routes

import (
	"github.com/Jeboczek/TechnikInformatykBackend/cache"
	"github.com/gin-gonic/gin"
)

func DatabaseVersion(c *gin.Context) {
	cacheInstance := cache.GetCacheInstance()
	c.JSON(200, cacheInstance.GetDatabaseVersion())
}
