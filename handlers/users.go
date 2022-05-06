package handlers

import (
	"fmt"
	"golang-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// /* Delete user handler */
// func DeleteUser(ctx *gin.Context) {
// 	entry := new(commons.LoggingEntry)
// 	payload := commons.LoggingPayload{Name: "DeleteUser", UID: "userID.(string)"}
// 	logger := commons.GetLoggerInstance(*commons.LogUsers)
// 	var user commons.User
// 	// Call BindJSON to bind the received JSON to user
// 	err := ctx.BindJSON(&user)
// 	if err != nil {
// 		err := fmt.Errorf("delete user: %v", err)
// 		payload.Message = err.Error()
// 		logger.Log(entry.New(payload, ctx, logging.Error, http.StatusBadRequest))
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		log.Println(err.Error())
// 		return
// 	}
// 	err = models.DeleteUser(ctx, user.Uid)
// 	if err != nil {
// 		err := fmt.Errorf("delete user: %v", err)
// 		payload.Message = err.Error()
// 		logger.Log(entry.New(payload, ctx, logging.Error, http.StatusInternalServerError))
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error":  err.Error(),
// 			"detail": map[string]interface{}{"userID": user.Uid},
// 		})
// 		log.Println(err.Error())
// 		return
// 	} else {
// 		log.Printf("User successfully deleted: %v\n", user.Uid)
// 		ctx.JSON(http.StatusCreated, gin.H{
// 			"success": "User successfully deleted",
// 			"detail":  map[string]interface{}{"userID": user.Uid},
// 		})
// 		payload.Message = "success"
// 		logger.Log(entry.New(payload, ctx, logging.Debug, http.StatusCreated))
// 	}
// }

// /* Get all users handler */
// func GetAllUsers(ctx *gin.Context) {
// 	entry := new(commons.LoggingEntry)
// 	payload := commons.LoggingPayload{Name: "GetAllUsers", UID: "userID.(string)"}
// 	logger := commons.GetLoggerInstance(*commons.LogUsers)

// 	users, err := models.GetAllUsers(ctx)
// 	if err != nil {
// 		err := fmt.Errorf("listing all users: %v", err)
// 		log.Println(err.Error())
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		payload.Message = err.Error()
// 		logger.Log(entry.New(payload, ctx, logging.Error, http.StatusInternalServerError))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"response": users})
// 	payload.Message = "success"
// 	logger.Log(entry.New(payload, ctx, logging.Debug, http.StatusOK))
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

// func DeleteUser(ctx *gin.Context) {

// 	userID := ctx.Param("user_id")
// 	if userID == "" {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
// 		return
// 	}

// 	if err := models.DeleteUser(userID); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"success": "User was successfully deleted"})
// }

// // GetUsers will return all the users
// func GetUsers(ctx *gin.Context) {
// 	if err := models.CheckUserType(ctx, "ADMIN"); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	// call GetUsers to get all users response
// 	response, err := models.GetUsers(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
// 		return
// 	}

// 	var allusers []bson.M
// 	if err = response.All(ctx, &allusers); err != nil {
// 		log.Println(err.Error())
// 	}
// 	// send the response message
// 	ctx.JSON(http.StatusOK, allusers[0])
// }
