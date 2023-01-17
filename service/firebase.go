package service

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

// Start Firebase Instance
func FirebaseInstance() *auth.Client {
	// Create a new client and connect to the firebase server
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalln("Error initializing the app :" + err.Error())
	}
	// Access auth service from the default app
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalln("Unable to establish connection :" + err.Error())
	}
	return client
}
