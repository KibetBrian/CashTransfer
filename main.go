package main

import (
	"github.com/KibetBrian/fisa/routes"
	"github.com/gin-gonic/gin"
)




func main (){
	router :=gin.Default();
	router.GET("/hello", routes.SayHello)
	router.POST("/user/insert", routes.RegisterUser)
	router.Run("localhost:8080")
}

