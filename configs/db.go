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
	envKeys = []string{"DB_HOST","DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	envValues = map[string]string{
		"DB_HOST": "",
        "DB_PORT": "",
        "DB_USER": "",
        "DB_PASSWORD": "",
        "DB_NAME": "",
	}
	Db          *gorm.DB
)


func FetchEnvValues(){
	for _, key := range envKeys {
        value, ok := os.LookupEnv(key);
		if !ok {
			log.Fatalf("Failed to get environment variable%v", key)
		}
		envValues[key] = value
	}
}

func ConnectDb() (*gorm.DB, error) {
	FetchEnvValues()

	//Converts string port from env to port number
	var _, err = strconv.Atoi(envValues["DB_PORT"])
	if err != nil {
		log.Fatalf("Error converting string port to int. Error: %v", err)
	}

	//Postgres Connection details
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", envValues["DB_HOST"], envValues["DB_PORT"], envValues["DB_USER"], envValues["DB_PASSWORD"], envValues["DB_NAME"])

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err !=nil {
		log.Fatalf("Database connection error: %v", err)
	}

	//Auto migrations
	db.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{})
	return db, nil
}
