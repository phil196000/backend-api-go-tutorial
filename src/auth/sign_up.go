package auth

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)



func SignUpWithEmailAndPassword(c *gin.Context){
	var userData User
	firestoreClient := c.MustGet("firestoreClient").(*firestore.Client)
	client := c.MustGet("authClient").(*auth.Client)
if err := c.ShouldBindJSON(&userData); err !=nil{
	c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
}
params := (&auth.UserToCreate{}).Email(userData.Email).Password(userData.Password).DisplayName(userData.FirstName + " " + userData.LastName)
user,err := client.CreateUser(context.Background(),params)
if err != nil {
	c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to sign up"})
	return
}
userDataInit := map[string]interface{}{
	"email":userData.Email,
	"firstName":userData.FirstName,
	"lastName":userData.LastName,
	"id":user.UID,
}
_,err = firestoreClient.Collection("users").Doc(user.UID).Set(context.Background(),userDataInit)
c.JSON(http.StatusOK,gin.H{"message":"User signed up successfully", "userID":user.UID})
}