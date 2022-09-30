package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/KibetBrian/fisa/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = os.Getenv("DB_HOST")
	port=os.Getenv("DB_PORT")
	user=os.Getenv("DB_USER")
	password=os.Getenv("DB_PASSWORD")
	database=os.Getenv("DB_DATABASE")
	Db          *gorm.DB
)

func checkError(err error) bool {
	if err != nil {
		return err != nil
	}
	return false
}

func ConnectDb() (*gorm.DB, error) {

	//Converts string port from env to port number
	var _, err = strconv.Atoi(port)
	if err != nil {
		log.Fatal("Error converting string port to int")
	}

	//Postgres Connection details
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if checkError(err) {
		return nil, err
	}
	//Auto migrations
	db.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{})
	return db, nil
}
