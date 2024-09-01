package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	db.InitDb()

	server := gin.Default()

	server.GET("/", getHome)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getSingleEvent)
	server.POST("/events", postEvent)

	server.Run(":8080")
}

func getHome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"asdf": "asdf"})
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println(fmt.Errorf("error: Could not fetch events. Try again later, details:%v", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully fetched all events", "event": *events})
}

func postEvent(ctx *gin.Context) {
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

func getSingleEvent(ctx *gin.Context) {
	eventIdStr := ctx.Param("id")

	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}

	fmt.Println("event id", eventId)

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Could not find the event with id: %v", eventId)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully fetched the event", "event": event})
}
