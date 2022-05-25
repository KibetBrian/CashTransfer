package configs

import (
	"fmt"
	"log"
	"strconv"

	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/utils"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	host, _ = utils.GetEnvVal("DB_HOST")
	port, _ = utils.GetEnvVal("DB_PORT")
	user, _  = utils.GetEnvVal("DB_USER")
	password, _ = utils.GetEnvVal("DB_PASSWORD")
	database, _ = utils.GetEnvVal("DB_DATABASE")
    Db *gorm.DB
)

func checkError (err error) bool{
	if err != nil {
		return err!=nil
	}
	return false
}


func ConnectDb () (*gorm.DB, error){

	//Converts string port from env to port number
	var _, err= strconv.Atoi(port);
	if err != nil {
		log.Fatal("Error converting string port to int")
	}
	
	//Postgres Connection details
	psqlInfo :=  fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,database)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	
	if checkError(err) {
		return nil, err
	}
	//Auto migrations
	db.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{})
	return db, nil;
}


