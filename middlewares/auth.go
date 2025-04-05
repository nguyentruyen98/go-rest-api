package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lolstate/lol-api-service/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(401, gin.H{"error": "Unauthorized"})
		context.Abort()
		return
	}
	context.Set("userId", userId)
	context.Next()
	// context.JSON(200, gin.H{"userId": userId})
}
