package api

import (
	"log"

	"github.com/KibetBrian/fisa/auth"
	"github.com/KibetBrian/fisa/utils"
)

func NewMaker() (auth.Maker) {
	secretKey, err := utils.GetEnvVal("JWT_SECRET_KEY")
	if err != nil {
		log.Fatal("Error occurred while retrieving env val. Err: ", err)
		return nil
	}

	maker, err := auth.NewMaker(secretKey)
	if err != nil {
		log.Fatal("Error occurred while creating new make. Err ", err)
		return nil
	}
	
	return maker
}