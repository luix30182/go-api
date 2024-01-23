package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"mario-mtz.com/rest-api/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)

}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cold not parse request"})
		return
	}

	event.ID = 1
	event.UserId = 1

	event.Save()

	context.JSON(http.StatusCreated, event)
}
