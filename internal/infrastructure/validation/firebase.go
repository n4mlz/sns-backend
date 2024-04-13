package validation

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

// initialize app with ServiceAccountKey.json
func InitFirebaseApp() error {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	FirebaseApp = app

	if err != nil {
		log.Printf("error initializing app: %v\n", err)
		return err
	}

	return nil
}

func VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	client, err := FirebaseApp.Auth(ctx)
	if err != nil {
		log.Printf("error getting Auth client: %v\n", err)
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		return nil, err
	}

	log.Printf("Verified ID token: %v\n", token)

	return token, nil
}
