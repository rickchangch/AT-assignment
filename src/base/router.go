package base

import (
	"github.com/gin-gonic/gin"
	"github.com/rickchangch/at-assignment/controller"
	"github.com/rickchangch/at-assignment/middleware"
)

func UrlMap(router *gin.Engine) {
	// Declare routers for auth endpoints
	auth := router.Group("/")
	{
		auth.POST("sign-up", controller.AuthController.SignUp)
		auth.POST("sign-in", controller.AuthController.SignIn)
	}

	// Declare routers for v1 REST APIs
	v1 := router.Group("/v1")
	v1.Use(middleware.AuthHandler)
	{
		v1.GET("/users", controller.UserController.ListUsers)
		v1.GET("/users/:id", controller.UserController.ListUsers)
		v1.PATCH("/users/:id", controller.UserController.ListUsers)
		v1.DELETE("/users/:id", controller.UserController.ListUsers)
		v1.POST("/users/search", controller.UserController.ListUsers)
	}
}
