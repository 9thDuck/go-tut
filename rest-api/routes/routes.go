package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", GetHome)
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetSingleEvent)

	server.POST("/users/signup", Signup)
	server.POST("/users/login", Login)

	authenticated := server.Group("/").Use(middlewares.Authenticate)
	authenticated.POST("/events", PostEvent)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)

}
