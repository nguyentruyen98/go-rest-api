package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lolstate/lol-api-service/models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch event by id"})
		return
	}
	event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't register for event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success", "event": event})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch event by id"})
		return
	}
	event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't cancel registration for event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success", "event": event})
}
