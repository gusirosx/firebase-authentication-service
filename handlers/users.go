package handlers

import (
	"firebase/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers will return a list of all users
func GetUsers(ctx *gin.Context) {
	values := ctx.Request.URL.Query()
	if len(values) != 0 {
		GetUserByParams(ctx)
		return
	}

	// call GetUsers to get all users response
	response, err := models.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		return
	}
	// send the response message
	ctx.JSON(http.StatusOK, response)
}

// GetUser will return a specific user by a given parameter
func GetUserByParams(ctx *gin.Context) {

	// call GetUser to get the user
	user, code, err := models.GetUserByParams(ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}
	// send the response message
	ctx.JSON(code, gin.H{
		"success": "User successfully retrieved",
		"user":    user,
	})
}

// GetUser will return a specific user
func GetUser(ctx *gin.Context) {

	// call GetUser to get the user
	user, code, err := models.GetUser(ctx)
	if err != nil {
		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}
	// send the response message
	ctx.JSON(code, gin.H{
		"success": "User successfully retrieved",
		"user":    user,
	})
}

// CreateUser create a user in the postgres database
func CreateUser(ctx *gin.Context) {
	// create an empty user of type entity.User
	var user models.User

	// decode the json request to user
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call CreateUser to create the user
	if err := models.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// send the response message
	ctx.JSON(http.StatusCreated, gin.H{"success": fmt.Sprintf("User %s was successfully created", user.DisplayName)})
}

func UpdateUser(ctx *gin.Context) {
	var user models.User

	userID := ctx.Param("id")
	if len(userID) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
		return
	}
	// Call BindJSON to bind the received JSON to user
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateUser(userID, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// send the response message
	ctx.JSON(http.StatusOK, gin.H{
		"success": "user was successfully updated",
		"user":    user})
}

func DeleteUser(ctx *gin.Context) {

	userID := ctx.Param("id")
	if len(userID) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
		return
	}

	if err := models.DeleteUser(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "User was successfully deleted"})
}
