package utils

import (
	"math/rand"
	"net/mail"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}
//Check user email input validity
func ValidateEmail(address string) bool {
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

