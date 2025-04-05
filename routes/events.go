package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lolstate/lol-api-service/models"
)

func getEvent(context *gin.Context) {
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
	context.JSON(http.StatusOK, gin.H{"message": "success", "event": event})
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}
func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the payload"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't create events"})
		return
	}
	context.JSON(http.StatusBadRequest, gin.H{"message": "Event created", "event": event})

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch event by Id"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "You are not allowed to update this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request data"})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't updated event by Id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch event by Id"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't delete event by Id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully", "event": event})

}
