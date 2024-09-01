package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	db.InitDb()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
