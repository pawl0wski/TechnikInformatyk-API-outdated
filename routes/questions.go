package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pawl0wski/TechnikInformatykBackend/cache"
)

func Questions(c *gin.Context) {
	cacheInstance := cache.GetCacheInstance()
	questions := cacheInstance.GetQuestions()
	c.JSON(200, questions)
	questions = nil
}
