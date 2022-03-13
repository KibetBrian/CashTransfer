package api

import (
	"log"

	"github.com/KibetBrian/fisa/auth"
	"github.com/KibetBrian/fisa/handlers"
	"github.com/KibetBrian/fisa/utils"
	"github.com/gin-gonic/gin"
)

var UserRoutes = func ( router *gin.Engine){
	router.GET("/hello", handlers.SayHello)
	router.POST("/user/login", handlers.Login)
	router.POST("/user/register", handlers.RegisterUser)
	secretKey, err := utils.GetEnvVal("JWT_SECRET_KEY")
	if err != nil {
		log.Fatal("Error occurred while retrieving env val. Err: ", err)
		return
	}
	maker, r := auth.NewMaker(secretKey)
	if err != nil {
		log.Fatal("Error occurred while creating new make. Err ", r)
		return
	}
	router.GET("/user/refreshtoken", auth.AuthMiddleware(maker), handlers.RefreshToken)
}
