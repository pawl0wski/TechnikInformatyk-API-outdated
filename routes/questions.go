package routes

import (
	"github.com/Jeboczek/TechnikInformatykBackend/cache"
	"github.com/gin-gonic/gin"
)

func Questions(c *gin.Context) {
	cacheInstance := cache.GetCacheInstance()
	questions := cacheInstance.GetQuestions()
	c.JSON(200, questions)
}
