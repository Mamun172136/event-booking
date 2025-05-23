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

// func login(context *gin.Context){
// 	var user models.User

// 	err := context.ShouldBindBodyWithJSON(&user)

// 	if err != nil{
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
// 		return
// 	}

// 	err = user.ValidateCredentials()

// 	if err != nil{
// 		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticte user."})
// 		return
// 	}

// 	token,err :=utils.GenerateToken(user.Email, user.ID)

// 	if err != nil{
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate token."})
// 		return
// 	}
// 	context.JSON(http.StatusOK, gin.H{"message": "login successful!!", "token":token})
// }
func login(context *gin.Context) {
    var user models.User

    // Use ShouldBindJSON instead of ShouldBindBodyWithJSON for simpler cases
    if err := context.ShouldBindJSON(&user); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{
            "message": "Could not parse request data",
            "error":   err.Error(), // Include the actual error for debugging
        })
        return
    }

    // Validate required fields
    if user.Email == "" || user.Password == "" {
        context.JSON(http.StatusBadRequest, gin.H{
            "message": "Email and password are required",
        })
        return
    }

    // Validate credentials
    if err := user.ValidateCredentials(); err != nil {
        context.JSON(http.StatusUnauthorized, gin.H{
            "message": "Invalid email or password", // Generic message for security
        })
        return
    }

    // Generate token
    token, err := utils.GenerateToken(user.Email, user.ID)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "message": "Could not generate authentication token",
            "error":   err.Error(), // Log the actual error for debugging
        })
        return
    }

    // Successful login
    context.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "token":   token,
        "user_id": user.ID, // Include user ID for client reference
    })
}

