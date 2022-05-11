package models

import (
	"context"
	"fmt"
	"golang-jwt/service"
	"log"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

// Create an unexported global variable to hold the firebase connection pool.
var client *auth.Client = service.FirebaseInstance()

// Get one user from the DB by its id
func GetUser(ctx *gin.Context) (User, int, error) {
	var urec *auth.UserRecord
	// var verror error
	values := ctx.Request.URL.Query()

	ID, err := getID(ctx)
	if err == nil {
		urec, err = client.GetUser(ctx, ID) // get the user by it's ID
	} else if _, ok := values["email"]; ok { // get the user by it's email
		urec, err = client.GetUserByEmail(ctx, values["email"][0])
	} else if _, ok := values["phoneNumber"]; ok { // get the user by it's phone
		urec, err = client.GetUserByPhoneNumber(ctx, "+"+values["phoneNumber"][0])
	} else {
		err = fmt.Errorf("invalid search parameters")
	}
	if err != nil {
		return User{}, 400, err
	}
	// Assemble the payload for client response
	user := User{
		Uid:         urec.UID,
		Email:       urec.Email,
		PhoneNumber: urec.PhoneNumber,
		DisplayName: urec.DisplayName,
		PhotoURL:    urec.PhotoURL}

	return user, 200, nil
}

/* Get user by e-mail function */

/* Get user by phone function */

/* Get user by UID function */
// Get all users from firebase
func GetUsers(ctx *gin.Context) ([]User, error) {
	var users []User
	iter := client.Users(ctx, "")
	for {
		urec, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("error listing users:", err.Error())
			return users, err
		}
		user := User{
			Uid:         urec.UID,
			Email:       urec.Email,
			PhoneNumber: urec.PhoneNumber,
			DisplayName: urec.DisplayName,
			PhotoURL:    urec.PhotoURL,
		}
		users = append(users, user)
	}
	log.Println("Successfully fetched users data")
	return users, nil
}

// Create one user into Firebase
func CreateUser(user User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	params := (&auth.UserToCreate{}).
		Email(user.Email).
		PhoneNumber(user.PhoneNumber).
		Password(user.Password).
		DisplayName(user.DisplayName).
		PhotoURL(user.PhotoURL)

	_, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("unable to create user")
	}
	return nil
}

// Update one user from the DB by its id
func UpdateUser(id string, user User) error {

	params := (&auth.UserToUpdate{}).
		Email(user.Email).
		PhoneNumber(user.PhoneNumber).
		DisplayName(user.DisplayName).
		PhotoURL(user.PhotoURL)

	_, err := client.UpdateUser(context.Background(), id, params)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("unable to update user")
	}

	return nil
}

// Delete one user from the DB by its id
func DeleteUser(id string) error {
	// Call the DeleteUser method by passing a valid userID
	if err := client.DeleteUser(context.Background(), id); err != nil {
		log.Println(err.Error())
		return fmt.Errorf("unable to update user")
	}
	return nil
}
