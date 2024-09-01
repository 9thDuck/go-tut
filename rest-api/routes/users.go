package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	var user models.User

	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	emailValid := user.IsEmailValid()

	if !emailValid {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email address supplied is not valid"})
		return
	}

	if len(user.Password) < 8 || len(user.Password) > 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "password needs to be at least 8 characters long and at most 20 characters long"})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}
