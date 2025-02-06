package routes

import (
	"net/http"

	"example.com/models"
	"example.com/utils"
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

func login(context *gin.Context){
	var user models.User

	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticte user."})
		return
	}

	token,err :=utils.GenerateToken(user.Email, user.ID)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate token."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "login successful!!", "token":token})
}

