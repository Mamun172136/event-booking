package routes

import (
	"net/http"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context){
	var user models.User

	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"could not parse request data."})
		return
	}
	err= user.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"could not save user."})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"messge":"user created successfully."})
}

