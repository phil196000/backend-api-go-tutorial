package auth

import (
	"net/http"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func SignInWithEmailAndPassword(c *gin.Context){
	

	// Retrieve the email and password from the context
	email := c.GetString("email")
	password := c.GetString("password")

	// Validate the email and password as needed
	if email == "admin" && password == "password" {
		authClient.
	}else {
		// Authentication failed
		c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed"})
	}
}

func signInAndGenerateToken(email,password string)(string, error){
// clients
	authClient := c.MustGet("authClient").(auth.Client)
	firestoreClient := c.MustGet("firestoreClient").(firestore.Client)

params := (&auth.UserRecord{}).Email(email).Password(password)
	user, err := authClient.GetUserByEmail(context.Background(), email)
	if err != nil {
		log.Printf("User retrieval error: %v\n", err)
		return "", err
	}

	if user == nil {
		log.Println("User not found")
		return "", nil
	}

	// Retrieve additional user data from Firestore
	doc, err := firestoreClient.Collection("users").Doc(user.UID).Get(context.Background())
	if err != nil {
		log.Printf("Firestore retrieval error: %v\n", err)
		return "", err
	}

	// Extract email and roles from the Firestore document
	var userData struct {
		Email string   `firestore:"email"`
		Roles []string `firestore:"roles"`
	}
	err = doc.DataTo(&userData)
	if err != nil {
		log.Printf("Firestore data extraction error: %v\n", err)
		return "", err
	}

	// Set custom claims with roles
	customClaims := make(map[string]interface{})
	customClaims["roles"] = userData.Roles
	if err := authClient.SetCustomUserClaims(context.Background(), user.UID, customClaims); err != nil {
		log.Printf("Setting custom claims error: %v\n", err)
		return "", err
	}

	token, err := authClient.CustomToken(context.Background(), user.UID)
	if err != nil {
		log.Printf("Token generation error: %v\n", err)
		return "", err
	}

	return token, nil
}