package routes

import (
	"errors"
	"net/http"
	"vita-track-ai/models"
	"vita-track-ai/repository"

	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=8"`
	Name     string `form:"name" binding:"required"`
}

func signup(context *gin.Context) {
	var signupRequest SignupRequest
	err := context.ShouldBind(&signupRequest) //not with JSON as it will be a form data :)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to pass the values into the user object",
			"error":   err.Error(),
		})
		return
	}

	var user models.User
	user.Email = signupRequest.Email
	user.Password = &signupRequest.Password
	user.Name = signupRequest.Name

	_, err = repository.GetUserModelByEmail(user.Email)

	if err == nil {
		context.JSON(http.StatusConflict, gin.H{
			"message": "User Already Exists",
			"error":   errors.New("User Already Exists").Error(),
		})
		return
	}

	err = repository.SaveUser(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "There is some problem saving the user",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"users": user})
}

func login(context *gin.Context) {

}

func googleLogin(context *gin.Context) {

}
