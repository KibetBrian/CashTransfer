package api

import (
	"github.com/KibetBrian/fisa/auth"
	"github.com/KibetBrian/fisa/handlers"
	"github.com/gin-gonic/gin"
)

var UserRoutes = func ( router *gin.Engine){
	router.GET("/hello", handlers.SayHello)
	router.POST("/user/login", handlers.Login)
	router.POST("/user/register", handlers.RegisterUser)
	router.GET("/user/refreshtoken", auth.AuthMiddleware(NewMaker()), handlers.RefreshToken)
}
