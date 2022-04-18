package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpawlowskii/TechnikInformatykBackend/cache"
)

func Exams(c *gin.Context) {
	cacheInsance := cache.GetCacheInstance()
	exams := cacheInsance.GetExams()
	c.JSON(200, exams)
	exams = nil
}
