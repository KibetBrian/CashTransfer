package api

import(
	"github.com/gin-gonic/gin"
	"github.com/KibetBrian/fisa/handlers"
)

var AcccountRoutes = func(router *gin.Engine) {
	router.POST("/account/create", handlers.CreateAccount)
	router.GET("/account/delete", handlers.DeleteAccount)
	router.GET("/account/account", handlers.GetAccount)
}
