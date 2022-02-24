package services

import (
	"log"

	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	uuid "github.com/satori/go.uuid"
)

//Takes account id as input and return account object from the database
func GetAccount(id uuid.UUID) (bool, *models.Account){
	var account models.Account
	db, err := configs.ConnectDb()
	if err != nil{
		log.Fatal("Database connection error: ", err)
	}
	res := db.Where("id=?", id).First(&account)
	if res.RowsAffected < 1{
		return false, nil
	}
	return true, &account
}