package middlewares

import (
	"errors"
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(token string) (int64, error) {
	JWT_SECRET, err := utils.GetEnvVariable("JWT_SECRET")

	if err != nil {
		return 0, err
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return 0, errors.New("unexpected signing method")
		}

		return []byte(JWT_SECRET), nil
	})

	if err != nil || !parsedToken.Valid {
		return 0, errors.New("Unauthorized")
	}
	claims, typeMatch := parsedToken.Claims.(jwt.MapClaims)

	if !typeMatch {
		return 0, errors.New("Unauthorized")
	}

	userId := int64(claims["user_id"].(float64))

	return userId, nil
}

func SendUnauthorizedRes(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
}

func Authenticate(ctx *gin.Context) {
	authToken, err := ctx.Request.Cookie("access_token")

	if err != nil {
		SendUnauthorizedRes(ctx)
		return
	}

	userId, err := VerifyToken(authToken.Value)

	if err != nil {
		SendUnauthorizedRes(ctx)
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
