package routes

import (
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Welcome to the Event Booking API",
        
        })
    })
	server.GET("/events",getEvents)
	server.GET("/events/:id",getEvent)

	authenticated:= server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register",registerForEvent )
	authenticated.DELETE("/events/:id/register", )
	// server.POST("/events", middlewares.Authenticate, createEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}