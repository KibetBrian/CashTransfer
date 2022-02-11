package routes

import (
	"github.com/KibetBrian/fisa/controllers"
	"github.com/gin-gonic/gin"
)

var TransactionRoutes = func (router *gin.Engine){
	router.POST("/transaction/deposit", controllers.Deposit)
	router.POST("/transaction/send", controllers.Send)
}