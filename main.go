package main

import (
	"net/http"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func main (){
	server:=gin.Default()
	server.GET("/events",getEvents)
	server.POST("/events",createEvent)
	server.Run(":5000")
}

func getEvents(context *gin.Context){
	events:=models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context){
	var event models.Event
	err :=context.ShouldBindBodyWithJSON(&event)

	if err !=nil{
		context.JSON(http.StatusBadGateway, gin.H{"message":"could not parse request data."})
	}

	event.ID=1
	event.UserID=1
	context.JSON(http.StatusAccepted, gin.H{"message": "event created", "event":event})
}