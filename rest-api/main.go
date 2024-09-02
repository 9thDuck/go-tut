package main

import (
	"fmt"
	"maps"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"example.com/rest-api/constants"
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"example.com/rest-api/utils"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	missingEnvVariables := utils.VerifyEnvAndGetMisingVars(slices.Collect(maps.Keys(constants.ENV_VARIABLE_NAMES)))

	if len(missingEnvVariables) > 0 {
		panic(fmt.Errorf("missing .env variables: %v", missingEnvVariables))
	}

	db.InitDb()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
