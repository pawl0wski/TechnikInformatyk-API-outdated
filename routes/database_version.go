package routes

import (
	"github.com/gin-gonic/gin"
	dblocker "github.com/pawl0wski/technikinformatyk-backend/db_locker"
)

func DatabaseVersion(c *gin.Context) {
	dbLocker := dblocker.GetDBLockerInstance()
	c.JSON(200, dbLocker.GetDatabaseVersion())
}
