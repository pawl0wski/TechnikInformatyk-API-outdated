package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pawl0wski/technikinformatyk-backend/cache"
)

func Exams(c *gin.Context) {
	cacheInsance := cache.GetCacheInstance()
	exams := cacheInsance.GetExams()
	c.JSON(200, exams)
	exams = nil
}
