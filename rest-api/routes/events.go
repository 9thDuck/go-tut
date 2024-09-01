package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetHome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"asdf": "asdf"})
}

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println(fmt.Errorf("error: Could not fetch events. Try again later, details:%v", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully fetched all events", "event": *events})
}

func PostEvent(ctx *gin.Context) {
	var event models.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data.", "err": err, "time": time.Now()})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func GetSingleEvent(ctx *gin.Context) {
	eventIdStr := ctx.Param("id")
	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not find the event with id: %v", eventId)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully fetched the event", "event": event})
}

func UpdateEvent(ctx *gin.Context) {
	eventIdStr := ctx.Param("id")

	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}

	var event models.Event

	err = ctx.ShouldBindJSON(&event)

	event.ID = eventId

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}

	err = event.Update()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not update the event with id:%v", eventId)})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("Successfully updated the event with id:%v", eventId)})
}

func DeleteEvent(ctx *gin.Context) {
	eventIdStr := ctx.Param("id")

	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}

	err = models.DeleteEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not delete the event by id:%v", eventId)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Successfully deleted the event with id:%v", eventId)})
}
