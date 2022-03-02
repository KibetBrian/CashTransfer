package main

import (
	"log"

	"github.com/KibetBrian/fisa/api"
	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/gin-gonic/gin"
)

func main() {
	
	 server := &api.Server{
		 Router: gin.Default(),
	 }
	 server.Serve("localhost:8080")
	
	db, err := configs.ConnectDb() 
	if err != nil{
		log.Fatal("Database connection failed. Err: ",err)
	}
	db.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{})

}
