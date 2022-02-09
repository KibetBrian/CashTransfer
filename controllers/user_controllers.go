package controllers

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SayHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Hello"})
}

//Validate clientside email input
func validateEmail(address string) bool {
	_, err := mail.ParseAddress(address)
	return err == nil
}

//Hash user plain text password
func hashPassword(plainText string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//Function to compare hashed password
func compareHash(plainTextPassword, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainTextPassword))
	return err == nil
}

func RegisterUser(c *gin.Context) {

	var user models.User
	c.ShouldBindJSON(&user)
	id := uuid.New()
	user.Id = id
	db, err := configs.ConnectDb()
	DB = db
	if err != nil {
		fmt.Print("Error Connecting to the database")
	}

	isValid := validateEmail(user.UserEmail)
	if !isValid {
		c.JSON(400, "Invalid email")
		return
	}

	//Check if email already exists
	res := db.Where("user_email= ?", user.UserEmail).First(&user)
	if res.RowsAffected > 0 {
		c.JSON(409, gin.H{"Message": res.Error})
		return
	}
	user.Password, _ = hashPassword(user.Password)

	//Insert user into db
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": result.Error})
		return
	}

	c.JSON(200, "User Registered")
}

func Login(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)

	db, err := configs.ConnectDb()
	if err != nil {
		return
	}
	if err := validateEmail(user.UserEmail); err != true {
		c.JSON(403, "Invalid email")
		return
	}

	plainText := user.Password

	res := db.Where("user_email= ?", user.UserEmail).First(&user)
	if res.RowsAffected < 1 {
		c.JSON(404, "Email not found")
		return
	}

	isValid := compareHash(plainText, user.Password)
	if !isValid {
		c.JSON(403, "Invalid password")
		return
	}
	c.JSON(200, "Login Successful")

}
