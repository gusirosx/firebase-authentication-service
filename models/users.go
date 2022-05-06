package models

import (
	"context"
	"fmt"
	"golang-jwt/service"
	"log"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

// Create an unexported global variable to hold the firebase connection pool.
var client *auth.Client = service.FirebaseInstance()

// Get one user from the DB by its id
func GetUser(ctx *gin.Context) (User, int, error) {
	var urec *auth.UserRecord
	var err error
	values := ctx.Request.URL.Query()

	// get the userID from the ctx params, key is "id"
	if userID := ctx.Param("id"); userID == "" {
		err = fmt.Errorf("no user ID was provided")
	} else {
		urec, err = client.GetUser(ctx, userID)
	}

	if _, ok := values["email"]; ok {
		urec, err = client.GetUserByEmail(ctx, values["email"][0])

	} else if _, ok := values["phoneNumber"]; ok {
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

// /* Update user function */
// func (server *UserManagementServer) UpdateUser(ctx context.Context, u *pb.User) (*pb.User, error) {
// 	params := (&auth.UserToUpdate{}).
// 		Email(u.Email).
// 		PhoneNumber(u.Phone).
// 		DisplayName(u.DisplayName).
// 		PhotoURL(u.PhotoURL)

// 	urec, err := server.client.UpdateUser(ctx, u.Uid, params)
// 	if err != nil {
// 		log.Println("error updating user: ", err.Error())
// 		return nil, err
// 	}
// 	log.Println("Successfully updating user: ", urec.UID)
// 	// Assemble the payload for client response
// 	user := &pb.User{
// 		Uid:         urec.UID,
// 		Email:       urec.Email,
// 		Phone:       urec.PhoneNumber,
// 		DisplayName: urec.DisplayName,
// 		PhotoURL:    urec.PhotoURL}
// 	return user, nil
// }

// /* Delete user function */
// func (server *UserManagementServer) DeleteUser(ctx context.Context, u *pb.UserID) (*pb.Empty, error) {
// 	err := server.client.DeleteUser(ctx, u.Uid)
// 	if err != nil {
// 		log.Println("error deleting user: ", err.Error())
// 		return nil, err
// 	}
// 	log.Println("Successfully deleting user: ", u.Uid)
// 	return &pb.Empty{}, nil
// }

// /* List Users function */
// func (server *UserManagementServer) GetUsers(ctx context.Context, u *pb.Empty) (*pb.UserList, error) {
// 	var users *pb.UserList = &pb.UserList{}
// 	iter := server.client.Users(ctx, "")
// 	for {
// 		u, err := iter.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			log.Println("error listing users:", err.Error())
// 			return users, err
// 		}
// 		user := pb.User{
// 			Uid:         u.UID,
// 			Email:       u.Email,
// 			Phone:       u.PhoneNumber,
// 			DisplayName: u.DisplayName,
// 			PhotoURL:    u.PhotoURL,
// 		}
// 		users.Users = append(users.Users, &user)
// 	}
// 	log.Println("Successfully fetched users data")
// 	return users, nil
// }

// // Get all users from the DB by its id
// func GetUsers(ctx *gin.Context) (response *mongo.Cursor, err error) {
// 	recordPerPage, err := strconv.Atoi(ctx.Query("recordPerPage"))
// 	if err != nil || recordPerPage < 1 {
// 		recordPerPage = 10
// 	}

// 	startIndex, _ := strconv.Atoi(ctx.Query("startIndex"))
// 	matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
// 	groupStage := bson.D{{Key: "$group", Value: bson.D{
// 		{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
// 		{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
// 		{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}}}}}
// 	projectStage := bson.D{
// 		{Key: "$project", Value: bson.D{
// 			{Key: "_id", Value: 0},
// 			{Key: "total_count", Value: 1},
// 			{Key: "user_items", Value: bson.D{
// 				{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}}}}}}}

// 	var queryCtx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	response, err = collection.Aggregate(queryCtx, mongo.Pipeline{matchStage, groupStage, projectStage})
// 	if err != nil {
// 		log.Println(err.Error())
// 		return
// 	}
// 	return
// }

// // Update one user from the DB by its id
// func UpdateUser(id string, user entity.User) error {

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()
// 	// Declare a primitive ObjectID from a hexadecimal string
// 	idPrimitive, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return err
// 	}

// 	recoveredUser, err := GetUser(id)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return fmt.Errorf("error during user recovering")
// 	}
// 	// e-mail check
// 	if !strings.EqualFold(*user.Email, *recoveredUser.Email) {
// 		if err := emailVerify(ctx, *user.Email); err != nil {
// 			return err
// 		}
// 	}
// 	// phone check
// 	if !strings.EqualFold(*user.Phone, *recoveredUser.Phone) {
// 		if err := phoneVerify(ctx, *user.Phone); err != nil {
// 			return err
// 		}
// 	}

// 	var updatedUser primitive.D
// 	password := HashPassword(*user.Password)
// 	time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

// 	// Update user info with the new token's
// 	token, refreshToken, err := GenerateAllTokens(user)
// 	if err != nil {
// 		return fmt.Errorf("unable to generate the user token's")
// 	}

// 	updatedUser = append(updatedUser,
// 		bson.E{Key: "userName", Value: user.UserName},
// 		bson.E{Key: "firstName", Value: user.FirstName},
// 		bson.E{Key: "lastName", Value: user.LastName},
// 		bson.E{Key: "password", Value: &password},
// 		bson.E{Key: "email", Value: user.Email},
// 		bson.E{Key: "phone", Value: user.Phone},
// 		bson.E{Key: "userType", Value: user.UserType},
// 		bson.E{Key: "token", Value: token},
// 		bson.E{Key: "refreshtoken", Value: refreshToken},
// 		bson.E{Key: "updated", Value: time})
// 	opt := options.Update().SetUpsert(true)
// 	update := bson.D{{Key: "$set", Value: updatedUser}}
// 	_, err = collection.UpdateByID(ctx, idPrimitive, update, opt)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return fmt.Errorf("unable to update user")
// 	}

// 	// filter := bson.M{"uid": user.UID}
// 	// update := bson.D{{Key: "$set", Value: updateToken}}
// 	// _, err = userCollection.UpdateOne(ctx, filter, update, opt)
// 	// if err != nil {
// 	// 	return fmt.Errorf("unable to update the user token's")
// 	// }

// 	return nil
// }

// // Delete one user from the DB by its id
// func DeleteUser(id string) error {
// 	var queryCtx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	// Declare a primitive ObjectID from a hexadecimal string
// 	idPrimitive, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return err
// 	}

// 	// Call the DeleteOne() method by passing BSON
// 	res, err := collection.DeleteOne(queryCtx, bson.M{"_id": idPrimitive})
// 	if err != nil {
// 		log.Println(err.Error())
// 		return fmt.Errorf("unable to delete user")
// 	} else if res.DeletedCount == 0 {
// 		return fmt.Errorf("there is no such user for be deleted")
// 	}
// 	return nil
// }
