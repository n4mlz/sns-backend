package validation

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type FirebaseApp struct {
	*firebase.App
}

// initialize app with ServiceAccountKey.json
func NewFirebaseApp() (*FirebaseApp, error) {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing app: %v\n", err)
		return nil, err
	}

	return &FirebaseApp{app}, nil
}

func (app *FirebaseApp) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	client, err := app.Auth(ctx)
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
