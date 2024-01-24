package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/words", getWords)
	server.GET("/words/:id", getWord)
	server.POST("words", createWord)
	server.PUT("/words/:id", updateWord)
	server.DELETE("/words/:id", deleteWord)
}
