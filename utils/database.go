package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	_ "strconv"
	"github.com/joho/godotenv"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
)
var (
	host = GetEnvValue("DB_HOST")
	port = GetEnvValue("DB_PORT")
	user  = GetEnvValue("DB_USER")
	password = GetEnvValue("DB_PASSWORD")
	database = GetEnvValue("DB_DATABASE")
)

func GetEnvValue (key string) string {
	err := godotenv.Load();
	if err != nil {
		log.Fatal ("Error loading .env file")
	}
	return os.Getenv(key);
}

func main (){

	var intPort, err= strconv.Atoi(GetEnvValue("DB_PORT"));
	if err != nil {
		log.Fatal("Failed to convert port to string")
	}

	psqlInfo :=  fmt.Sprintf("host =%s, port=%d, user=%s, password=%s, database=%s, sslmode=disable", host,intPort, user, password, database)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err =db.Ping();
	if err != nil {
		panic(err)
	}
	fmt.Print("Database Connected")

	defer db.Close()
	fmt.Print(GetEnvValue("DB_HOST"))
}

