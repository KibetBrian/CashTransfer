package api

import (
	"log"
	"os"

	"github.com/KibetBrian/fisa/auth"
)

func NewMaker() (auth.Maker) {
	secretKey, isFound:= os.LookupEnv("JWT_SECRET_KEY")
	if !isFound  {
		log.Fatal("Error occurred while retrieving env val. Key: ", "JWT_SECRET_KEY")
		return nil
	}

	maker, err := auth.NewMaker(secretKey)
	if err != nil {
		log.Fatal("Error occurred while creating new make. Err ", err)
		return nil
	}
	
	return maker
}