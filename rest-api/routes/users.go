package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
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
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func Login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the user credentials. valid email and password are required"})
		return
	}

	userPointer, err := user.ValidateCredentialsAndGetUser()

	user = *userPointer

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Given email and/or password are invalid"})
		return
	}

	token, err := user.GenerateToken()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong while logging in. Try again later."})
		return
	}

	maxAge, err := utils.GetJwtDuration()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong while logging in. Try again later."})
		return
	}

	ctx.SetCookie("access_token", token, int(maxAge), "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully logged in"})
}
