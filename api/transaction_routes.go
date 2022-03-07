package api

import (
	"github.com/KibetBrian/fisa/handlers"
	"github.com/gin-gonic/gin"
)

var TransactionRoutes = func (router *gin.Engine){
	router.POST("/transaction/deposit", handlers.Deposit)
	router.POST("/transaction/send", handlers.Send)
}
