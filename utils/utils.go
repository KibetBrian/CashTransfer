package utils

import (
	"fmt"
	"log"
	"math/rand"
	"net/mail"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

func init(){
	rand.Seed(time.Now().UnixNano())
	
	err := SetUpEnvironmentVariable(".env.example")
	if err != nil {
		log.Fatal(err)
	}
}

//Environment variable loader
func SetUpEnvironmentVariable(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Could not read environment variable: %v", err)
	}

	res := []string{}
	for _, v := range file{
		res = append(res, string(v))
	}
	
	resString := ""
	for _, v := range res{
		resString+=v
	}
	
	vals := strings.Split(resString, "\n")
	for _, v := range vals{
		v := strings.Split(v, "=")

		err := os.Setenv(v[0],v[1])
		if err!= nil {
			return fmt.Errorf("Could not set environment variable: %v", err)
		}

	}
	
	return nil
}

//Takes email address as an input and checks if it is valid
func CheckValidity(address string) bool {
	_, err := mail.ParseAddress(address)
	return err == nil
}

//Takes n as input and it return random number between 0 and n as long as n != 0
func GenerateRandInt(min,max int64) int64{
 	randomNumber := rand.Int63n(max-min)+min
	return randomNumber
}

//Takes interger n as input and returns random string of length n
func GenerateRandString (n int) string{
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	length := GenerateRandInt(0,int64(len(alphabet)))
	var sb strings.Builder
	for i:=0; i<=n; i++{
		sb.WriteByte(alphabet[GenerateRandInt(0,int64(length))])
	}

	return sb.String()
}

//Function to generate random username
func GenerateRandomUserName () string{
	const whiteSpace = " "

	var firstName strings.Builder
	var lastName strings.Builder
	var fullName strings.Builder

	//Generate random first and last names
	firstName.WriteString(GenerateRandString(int(GenerateRandInt(0,50))))
	lastName.WriteString(GenerateRandString(int(GenerateRandInt(0,50))))

	//Concatinate the names
	fullName.WriteString(firstName.String())
	fullName.WriteString(whiteSpace)
	fullName.WriteString(lastName.String())

	return fullName.String()
}

//Generate random password
func GenerateRandomPassword (min, max int64) string {
	const characters ="`1234567890qwertyuiopasdfghjklzxcvbnm,./!@#$%^&*()-=_+?><ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sb strings.Builder
	length := GenerateRandInt(min, max);

	for i:=0; i<=int(length); i++{
		sb.WriteByte(characters[GenerateRandInt(0,length)])
	}

	return sb.String()
}

//Generates random email
func GenerateRandomEmail () string{
	const at = "@"
	var sb strings.Builder
	domains := []string{".com",".net",".tech"}
	providers:=[]string{"gmail","outlook","yahoo","icloud"}
	name := GenerateRandString(int(GenerateRandInt(0,64)));
	domain:=domains[rand.Intn(len(domains))]
	provider := providers[rand.Intn(len(domains))]

	sb.WriteString(name)
	sb.WriteString(at)
	sb.WriteString(provider)
	sb.WriteString(domain)

	return sb.String()
}

//Takes error as an input and return gin.H object
func ErrResponse (err error) gin.H {
	return gin.H{"Error":err.Error()}
}

//Get env value
func GetEnvVal(key string) (string, error) {
	err := godotenv.Load(".env.example")
	if err != nil{
		return "", err
	}
	
	return os.Getenv(key), nil
}

//Takes uuid as input and returns converted string of it
func UUIDString(u uuid.UUID) string{
	return u.String()
}


