package routes

import (
	"github.com/Jeboczek/TechnikInformatykBackend/cache"
	"github.com/gin-gonic/gin"
)

func Exams(c *gin.Context) {
	cacheInsance := cache.GetCacheInstance()
	exams := cacheInsance.GetExams()
	c.JSON(200, exams)
}
