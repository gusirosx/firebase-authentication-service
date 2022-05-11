package handlers

import (
	"fmt"
	"golang-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers will return a list of all users
func GetUsers(ctx *gin.Context) {
	// call GetUsers to get all users response
	response, err := models.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		return
	}
	// send the response message
	ctx.JSON(http.StatusOK, response)
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

// /* Update user handler */
// func UpdateUser2(ctx *gin.Context) {

// 	var user commons.User
// 	// Call BindJSON to bind the received JSON to user
// 	err := ctx.BindJSON(&user)
// 	if err != nil {
// 		err := fmt.Errorf("update user: %v", err)
// 		payload.Message = err.Error()
// 		logger.Log(entry.New(payload, ctx, logging.Error, http.StatusBadRequest))
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		log.Println(err.Error())
// 		return
// 	}
// 	u, err := models.UpdateUser(ctx, user)
// 	if err != nil {
// 		err := fmt.Errorf("update user: %v", err)
// 		payload.Message = err.Error()
// 		logger.Log(entry.New(payload, ctx, logging.Error, http.StatusInternalServerError))
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 			"detail": map[string]interface{}{
// 				"email":       user.Email,
// 				"phoneNumber": user.PhoneNumber,
// 				"userName":    user.DisplayName},
// 		})
// 		log.Println(err.Error())
// 		return
// 	} else {
// 		fmt.Printf("User successfully updated: %v\n", u.Uid)
// 		ctx.JSON(http.StatusCreated, gin.H{
// 			"success": "User successfully updated",
// 			"detail": map[string]interface{}{
// 				"email":       u.Email,
// 				"phoneNumber": u.PhoneNumber,
// 				"userName":    u.DisplayName},
// 		})
// 		payload.Message = "success"
// 		logger.Log(entry.New(payload, ctx, logging.Debug, http.StatusOK))
// 	}
// }

//=======================================================================================================================
// var validate = validator.New()

// func UpdateUser(ctx *gin.Context) {
// 	var user entity.User

// 	userID := ctx.Param("user_id")
// 	if userID == "" {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
// 		return
// 	}

// 	if err := ctx.BindJSON(&user); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err := validate.Struct(user)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := models.UpdateUser(userID, user); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"success": "user was successfully updated"})
// }
