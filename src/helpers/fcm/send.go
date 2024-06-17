package fcm

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/v4/messaging"
	"github.com/eCanteens/backend-ecanteens/src/config"
)

func SendToToken(registrationToken string) {
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Notification Test",
			Body:  "Hello React!!",
		},
		Token: registrationToken,
	}

	response, err := config.FCM.Send(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}
