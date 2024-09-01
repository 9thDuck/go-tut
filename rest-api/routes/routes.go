package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", GetHome)
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetSingleEvent)
	server.POST("/events", PostEvent)
	server.PUT("/events/:id", UpdateEvent)
	server.DELETE("/events/:id", DeleteEvent)
}
