package routes

import (
	"github.com/Jeboczek/TechnikInformatykBackend/cache"
	"github.com/gin-gonic/gin"
)

func Image(c *gin.Context) {
	cacheInstance := cache.GetCacheInstance()
	c.Data(200, "image/png", cacheInstance.GetImage(c.Param("questionUuid")))
}
