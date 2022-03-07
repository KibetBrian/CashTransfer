package api

import (
	"github.com/KibetBrian/fisa/auth"
	"github.com/KibetBrian/fisa/handlers"
	"github.com/gin-gonic/gin"
)

var UserRoutes = func ( router *gin.Engine){
	router.GET("/hello", handlers.SayHello)
	
	maker := auth.JwtMaker{}
	auth := router.GET("/user").Use(auth.AuthMiddleware(&maker))
	auth.POST("/register", handlers.RegisterUser)
	auth.POST("/login", handlers.Login)
}
