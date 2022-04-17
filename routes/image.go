package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpawlowskii/TechnikInformatykBackend/cache"
)

func Image(c *gin.Context) {
	cacheInstance := cache.GetCacheInstance()
	c.Data(200, "image/png", cacheInstance.GetImage(c.Param("questionUuid")))
}
