package routes

import (
	"net/http"
	"strconv"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context){

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}

	event,err:=models.GetEventById(eventId)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
	}

	err = event.Register(userId)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register for  event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": " registered for  event"})

}