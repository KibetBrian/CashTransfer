package main

import (
	"log"

	"github.com/KibetBrian/fisa/api"
	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
)

func main() {

	api.RegisterRoutes()
	db, err := configs.ConnectDb()
	if err != nil {
		log.Fatal("Database connection failed. Err: ", err)
	}
	db.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{})

}
