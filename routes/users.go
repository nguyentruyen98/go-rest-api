package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lolstate/lol-api-service/models"
	"github.com/lolstate/lol-api-service/utils"
)

func signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the payload"})
		return
	}
	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})

}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the payload"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successfully", "token": token})

}
