package helpers

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

// FirebaseClient - fetch client
func FirebaseClient() (*db.Client, context.Context) {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: os.Getenv("DB_URL"),
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile(os.Getenv("SERVICE_ACCOUNT"))

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
	return client, ctx
}