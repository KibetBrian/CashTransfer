package api

import (
	"github.com/KibetBrian/fisa/handlers"
	"github.com/gin-gonic/gin"
)

var UserRoutes = func ( router *gin.Engine){
	router.GET("/hello", handlers.SayHello)
	router.POST("/user/login", handlers.Login)
	router.POST("/user/register", handlers.RegisterUser)
}
