// routes.go

package routes

import (
	authenticate "backend-api/src/auth"
	"middleware"

	"github.com/gin-gonic/gin"
	// "handler"
)

func SetupRoutes(router *gin.Engine) {

	// router.GET("/users", handler.User)
	router.POST("/signup",func(c *gin.Context){
		authenticate.SignUpWithEmailAndPassword(c)
	})
	router.Use(middleware.AuthorizationMiddleware)
	router.GET("signin", authenticate.SignInWithEmailAndPassword)
	// router.Use(middleware.AuthMiddleware())

	// router.POST("/users", createUser)
	// router.GET("/users/:id", getUserByID)
	// router.PUT("/users/:id", updateUser)
	// router.DELETE("/users/:id", deleteUser)
}
