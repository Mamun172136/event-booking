package routes

import (
	"database/sql"
	"net/http"
	"strconv"

	"example.com/db"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

// func registerForEvent(context *gin.Context){

// 	userId := context.GetInt64("userId")
// 	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

// 	if err != nil{
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
// 		return
// 	}

// 	event,err:=models.GetEventById(eventId)

// 	if err != nil{
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
// 	}

// 	err = event.Register(userId)

// 	if err != nil{
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register for  event"})
// 		return
// 	}

// 	context.JSON(http.StatusCreated, gin.H{"message": " registered for  event"})

// }

func registerForEvent(context *gin.Context) {
    userId := context.GetInt64("userId")
    if userId == 0 {
        context.JSON(http.StatusUnauthorized, gin.H{"message": "authentication required"})
        return
    }

    eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{
            "message": "invalid event ID",
            "error":   err.Error(),
        })
        return
    }

    // Verify event exists
    event, err := models.GetEventById(eventId)
    if err != nil {
        if err == sql.ErrNoRows {
            context.JSON(http.StatusNotFound, gin.H{"message": "event not found"})
        } else {
            context.JSON(http.StatusInternalServerError, gin.H{
                "message": "failed to fetch event",
                "error":   err.Error(),
            })
        }
        return
    }

    // Check if already registered
    var exists bool
    err = db.DB.QueryRow(`
        SELECT EXISTS(SELECT 1 FROM registrations 
        WHERE event_id = ? AND user_id = ?)`, 
        eventId, userId).Scan(&exists)
    
    if exists {
        context.JSON(http.StatusConflict, gin.H{"message": "already registered for this event"})
        return
    }

    // Register user
    if err := event.Register(userId); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "message": "registration failed",
            "error":   err.Error(), // Remove in production
        })
        return
    }

    context.JSON(http.StatusCreated, gin.H{
        "message": "successfully registered for event",
        "eventId": event.ID,
    })
}