package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const keyName = "JWT_SECRET_KEY"

func GenerateToken(username string, duration time.Duration) (string, error) {
	secretKey , err := GetEnvSecretKey()
	if err!= nil{
		return "", fmt.Errorf("env load failed")
	}

	maker := JwtMaker{secretKey}

	token, err := maker.CreateToken(username, duration)
	if err!= nil{
		return "", fmt.Errorf("token creation failed. Err: %v",err)
	}

	return token, nil
}

func GetEnvSecretKey() (string, error) {
	err := godotenv.Load()
	if err != nil{
		return "", err
	}
	
	return os.Getenv(keyName), nil
}

