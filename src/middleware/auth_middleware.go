// auth_middleware.go

package middleware

import (
	"context"
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	client :=  c.MustGet("authClient").(auth.Client)

		idToken := c.GetHeader("Authorization")
		if idToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			log.Printf("Firebase token verification error: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired authorization token"})
			return
		}

		// Set the authenticated user ID in the context for further processing
		c.Set("userID", token.UID)

		c.Next()
	}
}

func AuthorizationMiddleware(c *gin.Context){
	// Get the Authorization header value
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == ""{
		// No Authorization header provided
		c.Header("WWW-Authenticate",`Basic realm="Restricted"`)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Authentication failed"})
		return
	}

	// Check if the Authorization header starts with "Basic "
	if !strings.HasPrefix(authHeader, "Basic "){
		// Invalid Authorization header format
		c.Header("WWW-Authenticate",`Basic realm="Restricted"`)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Authentication failed"})
		return
	}

	// Decode the base64-encoded credentials
	credentials, err := base64.StdEncoding.DecodeString(authHeader[6:])
	if err !=nil {
		// Error decoding credentials
		c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Authentication failed"})
		return
	}

	// Extract the email and password from the credentials
	credentialsStr := string(credentials)
	credentialsArr := strings.Split(credentialsStr,":")
	if len(credentialsArr) != 2 {
		// Invalid credentials format
		c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Authentication failed"})
		return 
	}

	email := credentialsArr[0]
	password := credentialsArr[1]

	// Store the extracted email and password in the context
	c.Set("email", email)
	c.Set("password", password)

	c.Next() // Continue to the next middleware or route handler
}