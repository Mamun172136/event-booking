package routes

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could no fetch events, try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "could no parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could no fetch events, try again later."})
		return
	}
	context.JSON(http.StatusOK, event)
}

// func createEvent(context *gin.Context) {
	
// 	var event models.Event
// 	err := context.ShouldBindBodyWithJSON(&event)

// 	if err != nil {
// 		context.JSON(http.StatusBadGateway, gin.H{"message": "could not parse request data."})
// 	}

// 	userId:=context.GetInt64("userId")
// 	event.UserID=userId
// 	err = event.Save()
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event."})
// 	}
// 	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
// }
func createEvent(context *gin.Context) {
    var event models.Event
    
    // Use ShouldBindJSON instead of ShouldBindBodyWithJSON
    if err := context.ShouldBindJSON(&event); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{
            "message": "Invalid request data",
            "error":   err.Error(), // Include error details for debugging
        })
        return // Make sure to return after error
    }

    // Validate required fields
    if event.Name == "" || event.Description == "" || event.Location == "" || event.DateTime.IsZero() {
        context.JSON(http.StatusBadRequest, gin.H{
            "message": "All fields are required",
        })
        return
    }

    // Get user ID from context
    userId := context.GetInt64("userId")
    if userId == 0 {
        context.JSON(http.StatusUnauthorized, gin.H{
            "message": "Not authenticated",
        })
        return
    }
    event.UserID = userId

    // Save to database
    if err := event.Save(); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "message": "Failed to create event",
            "error":   err.Error(), // Log actual error for debugging
        })
        return
    }

    // Successful response
    context.JSON(http.StatusCreated, gin.H{
        "message": "Event created successfully",
        "event":   event,
    })
}

// func updateEvent(context *gin.Context){

// 	eventId, err:= strconv.ParseInt(context.Param("id"), 10, 64)
// 	if err != nil{
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
// 		return
// 	}

// 	userId := context.GetInt64("userId")
// 	event, err := models.GetEventById(eventId)

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event data."})
// 	}

// 	if event.UserID  != userId{
// 		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update event"})
// 		return
// 	}
	

// 	var updatedEvent models.Event
// 	err = context.ShouldBindBodyWithJSON(&updatedEvent)
// 	if err != nil {
// 		context.JSON(http.StatusBadGateway, gin.H{"message": "could not parse request data."})
// 	}

// 	updatedEvent.ID=eventId
// 	err = updatedEvent.Update()

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event data."})
// 	}

// 	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
// }

// func updateEvent(context *gin.Context) {
//     // Parse event ID
//     eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
//     if err != nil {
//         context.JSON(http.StatusBadRequest, gin.H{
//             "message": "Invalid event ID",
//             "error":   err.Error(),
//         })
//         return
//     }

//     // Get user ID from context
//     userId := context.GetInt64("userId")
//     if userId == 0 {
//         context.JSON(http.StatusUnauthorized, gin.H{
//             "message": "Not authenticated",
//         })
//         return
//     }

//     // Verify event exists and belongs to user
//     event, err := models.GetEventById(eventId)
//     if err != nil {
//         if err == sql.ErrNoRows {
//             context.JSON(http.StatusNotFound, gin.H{
//                 "message": "Event not found",
//             })
//         } else {
//             context.JSON(http.StatusInternalServerError, gin.H{
//                 "message": "Failed to fetch event",
//                 "error":   err.Error(),
//             })
//         }
//         return
//     }

//     if event.UserID != userId {
//         context.JSON(http.StatusUnauthorized, gin.H{
//             "message": "Not authorized to update this event",
//         })
//         return
//     }

//     // Parse update data
//     var updatedEvent models.Event
//     if err := context.ShouldBindJSON(&updatedEvent); err != nil {
//         context.JSON(http.StatusBadRequest, gin.H{
//             "message": "Invalid request data",
//             "error":   err.Error(),
//         })
//         return
//     }

//     // Update event
//     updatedEvent.ID = eventId
//     if err := updatedEvent.Update(); err != nil {
//         context.JSON(http.StatusInternalServerError, gin.H{
//             "message": "Failed to update event",
//             "error":   err.Error(),
//         })
//         return
//     }

//     context.JSON(http.StatusOK, gin.H{
//         "message": "Event updated successfully",
//         "event":   updatedEvent,
//     })
// }

func updateEvent(context *gin.Context) {
    // 1. Get and validate ID parameter
    eventIdStr := context.Param("id")
    if eventIdStr == "" {
        context.JSON(http.StatusBadRequest, gin.H{
            "message": "Event ID is required",
            "details": "Missing ID in URL path",
        })
        return
    }

    eventId, err := strconv.ParseInt(eventIdStr, 10, 64)
    if err != nil || eventId <= 0 {
        context.JSON(http.StatusBadRequest, gin.H{
            "message": "Invalid event ID",
            "details": "ID must be a positive integer",
            "input":   eventIdStr,
        })
        return
    }

    // 2. Authentication check
    userId := context.GetInt64("userId")
    if userId == 0 {
        context.JSON(http.StatusUnauthorized, gin.H{
            "message": "Authentication required",
        })
        return
    }

    // 3. Verify event ownership
    event, err := models.GetEventById(eventId)
    if err != nil {
        if err == sql.ErrNoRows {
            context.JSON(http.StatusNotFound, gin.H{
                "message": "Event not found",
            })
        } else {
            context.JSON(http.StatusInternalServerError, gin.H{
                "message": "Failed to fetch event",
                "error":   err.Error(),
            })
        }
        return
    }

    if event.UserID != userId {
        context.JSON(http.StatusForbidden, gin.H{
            "message": "Not authorized to update this event",
        })
        return
    }

    // 4. Parse and validate update data
    var updateData struct {
        Name        string    `json:"name" binding:"required"`
        Description string    `json:"description" binding:"required"`
        Location    string    `json:"location" binding:"required"`
        DateTime    time.Time `json:"dateTime" binding:"required"`
    }

    if err := context.ShouldBindJSON(&updateData); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{
            "message": "Invalid request data",
            "error":   err.Error(),
        })
        return
    }

    // 5. Perform update
    updatedEvent := models.Event{
        ID:          eventId,
        Name:        updateData.Name,
        Description: updateData.Description,
        Location:    updateData.Location,
        DateTime:    updateData.DateTime,
        UserID:      userId,
    }

    if err := updatedEvent.Update(); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "message": "Failed to update event",
            "error":   err.Error(),
        })
        return
    }

    // 6. Return updated event
    context.JSON(http.StatusOK, gin.H{
        "message": "Event updated successfully",
        "event":   updatedEvent,
    })
}

// func deleteEvent(context *gin.Context){
// 	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
// 	if err != nil{
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
// 		return
// 	}

// 	event, err := models.GetEventById(eventId)

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event data."})
// 	}

// 	err =event.Delete()

// 	if err != nil{
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event data."})
// 	}
// 	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
// }

func deleteEvent(context *gin.Context) {
    // 1. Parse event ID
    eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
        return
    }

    // 2. Check authentication
    userId := context.GetInt64("userId")
    if userId == 0 {
        context.JSON(http.StatusUnauthorized, gin.H{"message": "not authenticated"})
        return
    }

    // 3. Verify event exists and belongs to user
    event, err := models.GetEventById(eventId)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event data."})
        return
    }

    if event.UserID != userId {
        context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete event"})
        return
    }

    // 4. Perform deletion
    if err := event.Delete(); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event."})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}