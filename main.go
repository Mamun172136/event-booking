package main

import (
	// "net/http"
	// "strconv"

	"example.com/db"
	// "example.com/models"
	"example.com/routes"
	"github.com/gin-gonic/gin"
)

func main (){
	db.InitDB()
	server:=gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":5000")
}

