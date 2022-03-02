package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes (){
	router := gin.Default()
	AcccountRoutes(router)
	TransactionRoutes(router)
	UserRoutes(router)
	router.Run()
}