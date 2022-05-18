package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pawl0wski/technikinformatyk-backend/routes/GET"
)

func SetupRoutes(server *gin.Engine) {
	server.GET("/api/exams", GET.Exams)
	server.GET("/api/questions", GET.Questions)
	server.GET("/api/image/:questionUuid", GET.Image)
	server.GET("/api/databaseVersion", GET.DatabaseVersion)
	server.GET("/api/ping", GET.Ping)
	server.NoRoute(GET.NotFound)
}
