package utils

import (
	"fmt"
	"os"

	"github.com/KibetBrian/fisa/auth"
)

func GetPayload (token string) (*auth.Payload, error){
	secretKey, isFound:= os.LookupEnv("JWT_SECRET_KEY")
	if !isFound{
		return nil, fmt.Errorf("Envirironment value not found. Key: JWT_SECRET_KEY")
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