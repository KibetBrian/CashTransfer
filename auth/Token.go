package auth

import (
	"fmt"
	"os"
	"time"
)

const keyName = "JWT_SECRET_KEY"

func GenerateToken(username string, duration time.Duration) (string, error) {
	secretKey, isFound := os.LookupEnv(keyName)
	if !isFound {
		return "", fmt.Errorf("Environment variable not found. Key: %s", keyName)
	}

	maker := JwtMaker{secretKey}

	token, err := maker.CreateToken(username, duration)
	if err!= nil{
		return "", fmt.Errorf("token creation failed. Err: %v",err)
	}

	return token, nil
}


