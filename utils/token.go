package utils

import (
	"fmt"

	"github.com/KibetBrian/fisa/auth"
)

func GetPayload (token string) (*auth.Payload, error){
	secretKey, err := GetEnvVal("JWT_SECRET_KEY")
	if err != nil{
		return nil, fmt.Errorf("env value retrieval failed. err: %v",err)
	}
	maker, err := auth.NewMaker(secretKey)
	if err != nil{
		return nil, fmt.Errorf("new make creation failed. err: %v",err)
	}
	payload, err := maker.VerifyToken(token)
	if err!= nil{
		return nil, err
	}
	return payload, nil
}