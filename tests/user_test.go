package tests

import (
	"log"
	"testing"

	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
)


func TestCreateAccount(t *testing.T){
	testUser := &models.User{
		Name: "Brian Kibet",
		Email: "briankibet2010@gmail.com",
		Password: "justPassword",
	}
	db, err := configs.ConnectDb()
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
	res :=db.Create(&testUser)
	log.Println(res)
}