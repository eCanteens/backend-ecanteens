package config

import (
	"context"
	"encoding/base64"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"

	"google.golang.org/api/option"
)

var FCM *messaging.Client

func getDecodedFireBaseKey() ([]byte, error) {

	fireBaseAuthKey := os.Getenv("FIREBASE_AUTH_KEY")

	decodedKey, err := base64.StdEncoding.DecodeString(fireBaseAuthKey)
	if err != nil {
		return nil, err
	}

	return decodedKey, nil
}

func SetupFirebase() {
	decodedKey, err := getDecodedFireBaseKey()
	if err != nil {
		panic(err.Error())
	}

	opt := option.WithCredentialsJSON(decodedKey)

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	//Messaging client
	ctx := context.Background()

	client, err := app.Messaging(ctx)
	if err != nil {
		panic("error getting Messaging client: %v\n")
	}

	FCM = client
}
