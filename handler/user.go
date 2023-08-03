package handler

import (
	"golang/helper"
	"golang/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage )
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", nil )
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	// token, err := h.jwtService.GenerateToken()

	formatter := user.FormatUser(newUser, "tokentokentokentoken")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter )


	c.JSON(http.StatusOK, response)
}



func (h *userHandler) Login(c *gin.Context){

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage )
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	LoggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage )
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(LoggedInUser, "tokentokentokentoken")
	response := helper.APIResponse("Login Successfully", http.StatusOK, "success", formatter )

	c.JSON(http.StatusOK, response)
	
}

func (h *userHandler) CheckEmailAvailibility(c *gin.Context){
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage )
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage )
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": IsEmailAvailable,
	}

	metaMessage := "Email Has Been Registered"

	if IsEmailAvailable {
		metaMessage = "Email is Available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data )
	c.JSON(http.StatusOK, response)
}