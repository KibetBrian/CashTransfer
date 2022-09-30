package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//Takes input p  of type string and return hashed value of p if no error
func HashPassword(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		return "", fmt.Errorf("password hashing failed")
	}
	
	return string(bytes), nil
}

//Takes hashed value and p, then compares if hash(p) == hashed value
func CompareHashAndPassword(hash,p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	return err == nil
}
