package GET

import (
	"github.com/gin-gonic/gin"
	dblocker "github.com/pawl0wski/technikinformatyk-backend/db_locker"
)

func Questions(c *gin.Context) {
	dbLocker := dblocker.GetDBLockerInstance()
	questions, err := dbLocker.GetQuestions()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
	} else {
		c.JSON(200, questions)
	}
}
