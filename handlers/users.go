package handlers

import (
	"fmt"
	"golang-jwt/entity"
	"golang-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Uid         string `json:"uid"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	DisplayName string `json:"displayname"`
	PhotoURL    string `json:"photourl"`
}

// CreateUser create a user in the postgres database
func CreateUser(ctx *gin.Context) {
	// create an empty user of type entity.User
	var user entity.User2

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
	ctx.JSON(http.StatusCreated, gin.H{"success": fmt.Sprint("User %s was successfully created", user.DisplayName)})
}

/* Create user handler */

// /* Update user handler */
// func UpdateUser(ctx *gin.Context) {
// 	entry := new(commons.LoggingEntry)
// 	payload := commons.LoggingPayload{Name: "UpdateUser", UID: "userID.(string)"}
// 	logger := commons.GetLoggerInstance(*commons.LogUsers)
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

// /* Get users auxiliary fuction */
// func GetUsers(ctx *gin.Context) {
// 	values := ctx.Request.URL.Query()
// 	if len(values) == 0 {
// 		GetAllUsers(ctx)
// 	} else {
// 		GetUser(ctx)
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

// /* Get user handler */
// func GetUser(ctx *gin.Context) {
// 	entry := new(commons.LoggingEntry)
// 	payload := commons.LoggingPayload{Name: "GetUser", UID: "userID.(string)"}
// 	logger := commons.GetLoggerInstance(*commons.LogUsers)
// 	values := ctx.Request.URL.Query()
// 	var user commons.User
// 	var err error
// 	if _, ok := values["email"]; ok {
// 		user, err = models.GetUserByEmail(ctx, values["email"][0])
// 	} else if _, ok := values["phoneNumber"]; ok {
// 		user, err = models.GetUserByPhone(ctx, "+"+values["phoneNumber"][0])

// 	} else {
// 		err = fmt.Errorf("invalid search parameters")
// 		payload.Message = err.Error()
// 		logger.Log(entry.New(payload, ctx, logging.Error, http.StatusBadRequest))
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		log.Println(err.Error())
// 		return
// 	}
// 	//Checks for errors and inform the user
// 	if err != nil {
// 		payload.Message = err.Error()
// 		logger.Log(entry.New(payload, ctx, logging.Error, http.StatusBadRequest))
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	log.Printf("User successfully retrieved: %v", user)
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"success": "User successfully retrieved",
// 		"user":    user,
// 	})
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

// // GetUser will return a specific user
// func GetUser(ctx *gin.Context) {

// 	// get the userID from the ctx params, key is "id"
// 	userID := ctx.Param("id")
// 	if userID == "" {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
// 		return
// 	}

// 	if err := models.MatchUserTypeToUid(ctx, userID); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// call GetUser to get the user
// 	user, err := models.GetUser(userID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	// send the response message
// 	ctx.JSON(http.StatusOK, user)
// }
