package api

import (
	"github.com/KibetBrian/fisa/auth"
	"github.com/KibetBrian/fisa/handlers"
	"github.com/gin-gonic/gin"
)

var AcccountRoutes = func(router *gin.Engine) {
	router.POST("/account/create", handlers.CreateAccount)
	router.GET("/account/delete", handlers.DeleteAccount)
	router.GET("/account/account", handlers.GetAccount)
	router.GET("/account/balance", auth.AuthMiddleware(NewMaker()), handlers.CheckBalance)
}
