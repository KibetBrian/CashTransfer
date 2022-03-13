package api

import (
	"log"

	"github.com/KibetBrian/fisa/auth"
	"github.com/KibetBrian/fisa/handlers"
	"github.com/KibetBrian/fisa/utils"
	"github.com/gin-gonic/gin"
)

var TransactionRoutes = func(router *gin.Engine) {
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
	router.POST("/transaction/deposit", handlers.Deposit)
	router.POST("/transaction/send", auth.AuthMiddleware(maker), handlers.Send)
}
