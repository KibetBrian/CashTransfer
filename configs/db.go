package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/KibetBrian/fisa/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	host = GetEnvValue("DB_HOST")
	port = GetEnvValue("DB_PORT")
	user  = GetEnvValue("DB_USER")
	password = GetEnvValue("DB_PASSWORD")
	database = GetEnvValue("DB_DATABASE")
    Db *gorm.DB
)

func checkError (err error) bool{
	if err != nil {
		return err!=nil
	}
	return false
}

func GetEnvValue (key string) string {
	err := godotenv.Load("../.env");
	if err != nil {
		log.Fatal (err)
	}
	return os.Getenv(key);
}

func ConnectDb () (*gorm.DB, error){

	//Converts string port from env to port number
	var _, err= strconv.Atoi(port);
	if err != nil {
		log.Fatal("Error converting string port to int")
	}
	
	//Postres Connection details
	psqlInfo :=  fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,database)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	
	if checkError(err) {
		return nil, err
	}
	//Auto migrations
	db.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{})
	return db, nil;
}


