package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpawlowskii/TechnikInformatykBackend/cache"
)

func DatabaseVersion(c *gin.Context) {
	cacheInstance := cache.GetCacheInstance()
	c.JSON(200, cacheInstance.GetDatabaseVersion())
}
