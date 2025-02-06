package main

import (
	"net/http"
	"strconv"

	"example.com/db"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

func main (){
	db.InitDB()
	server:=gin.Default()
	server.GET("/events",getEvents)
	server.GET("/events/:id",getEvent)
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

func getEvent(context *gin.Context){
    eventId,err:= strconv.ParseInt(context.Param("id"),10, 64)	
	if err != nil{
		context.JSON(http.StatusBadGateway, gin.H{"message":"could no parse event id."})
		return
	}
	event,err:=models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"could no fetch events, try again later."})
		return
	}
	context.JSON(http.StatusOK, event)
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