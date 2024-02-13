package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mario-mtz.com/rest-api/models"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
	}
	context.JSON(http.StatusOK, events)

}

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse vent"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cold not fetch the event"})

	}

	context.JSON(http.StatusOK, event)

}

func CreateEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cold not parse request"})
		return
	}

	userId := context.GetInt64("userId")

	event.UserId = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cold not create the event"})

	}

	context.JSON(http.StatusCreated, event)
}

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse vent"})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event"})
		return
	}

	var updateEvent models.Event

	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cold not parse request"})
		return
	}

	updateEvent.ID = eventId

	err = updateEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse vent"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
