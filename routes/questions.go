package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pawl0wski/TechnikInformatykBackend/cache"
)

func Questions(c *gin.Context) {
	cacheInstance := cache.GetCacheInstance()
	questions, err := cacheInstance.GetQuestions()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
	} else {
		c.JSON(200, questions)
	}
	questions = nil
}
