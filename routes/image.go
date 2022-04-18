package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpawlowskii/TechnikInformatykBackend/cache"
)

func Image(c *gin.Context) {
	cacheInstance := cache.GetCacheInstance()
	image := cacheInstance.GetImage(c.Param("questionUuid"))

	if len(image) != 0 {
		c.Data(200, "image/png", image)
	} else {
		NotFound(c)
	}
}
