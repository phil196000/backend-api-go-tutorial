// main.go

package main

import (
	"backend-api/src/routes"
	"context"
	"log"

	"middleware"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("../agape-22bbc-firebase-adminsdk-6phjp-00e0102463.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Firebase initialization error: %v\n", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Firebase Auth initialization error: %v\n", err)
	}

	firestoreClient, err := app.Firestore(context.Background())
	if err!=nil {
		log.Fatalf("Firestore initialization error: %v\n",err)
	}

	router := gin.Default()

	router.Use(middleware.FirebaseMiddleware(authClient,firestoreClient))
	
	routes.SetupRoutes(router)


	router.Run(":8080")
}
