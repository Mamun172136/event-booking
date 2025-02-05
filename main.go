package main

import (
	"net/http"

	"example.com/db"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

func main (){
	db.InitDB()
	server:=gin.Default()
	server.GET("/events",getEvents)
	server.POST("/events",createEvent)
	server.Run(":5000")
}

func getEvents(context *gin.Context){
	events,err:=models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"could no fetch events, try again later."})
		return
	}
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
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"could not parse request data."})
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "event created", "event":event})
}