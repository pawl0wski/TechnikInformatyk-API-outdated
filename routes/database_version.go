package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pawl0wski/technikinformatyk-backend/cache"
)

func DatabaseVersion(c *gin.Context) {
	cacheInstance := cache.GetCacheInstance()
	c.JSON(200, cacheInstance.GetDatabaseVersion())
}
