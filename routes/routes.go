package routes

import (
	"fmt"
	"net/http"
	"net/mail"
	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func SayHello (c *gin.Context){
	c.IndentedJSON(http.StatusOK, gin.H{"Message":"Hello"})
}

//Validate clientside email input
func validateEmail (address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "Invalid email", false
	}
	return addr.Address, true
}
//Hash user plain text password
func hashPassword (plainText string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword ([]byte(plainText), 10);
	if err != nil{
		return "", err
	}
	return string(bytes), nil
}
//Function to compare hashed password
func compareHash(plainTextPassword, hash string) bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainTextPassword))
	return err==nil
}


func RegisterUser (c *gin.Context){
	 var user models.User;
	 c.ShouldBindJSON(&user)
	 db, err := utils.ConnectDb()
	 if err != nil {
		fmt.Print("Error Connecting to the database")
	 }
	 email, isValid:=validateEmail(user.Email)
	 if  !isValid {
		 c.JSON(400, "Invalid email")
	 }
	 user.Email=email;
	 //Check if email already exists
	 res := db.Where("email= ?", user.Email).First(&user)
	if res.RowsAffected < 1{
		c.JSON(409, gin.H{"Message": "Email already exist"})
	}
	 user.Password,_=hashPassword(user.Password)

	 //Insert user into db
	  result := db.Create(&user)
	  if result.Error != nil {
		  c.JSON(500, gin.H{"Error":result.Error})
	  }

	 c.JSON(200, result)
}
