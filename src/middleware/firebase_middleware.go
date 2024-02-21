package middleware

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func FirebaseMiddleware(authClient *auth.Client, firestoreClient *firestore.Client) gin.HandlerFunc{
	return func(c *gin.Context){
		c.Set("authClient",authClient)
		c.Set("firestoreClient",firestoreClient)
		c.Next()
	}
}