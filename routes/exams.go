package routes

import (
	"github.com/gin-gonic/gin"
	dblocker "github.com/pawl0wski/technikinformatyk-backend/db_locker"
)

func Exams(c *gin.Context) {
	dbLocker := dblocker.GetDBLockerInstance()
	exams := dbLocker.GetExams()
	c.JSON(200, exams)
}
