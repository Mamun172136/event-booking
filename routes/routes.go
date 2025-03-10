package routes

import (
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events",getEvents)
	server.GET("/events/:id",getEvent)

	authenticated:= server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	// server.POST("/events", middlewares.Authenticate, createEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}