package controllers

import (
	"fmt"
	"net/http"
	"net/mail"
	"time"
	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func RegisterUser(c *gin.Context) {
	
	var user models.User
	c.ShouldBindJSON(&user)
	id := uuid.New()
	user.UserId = id
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
	user.Password, _ = utils.HashPassword(user.Password)

	//Insert user into db
	db.AutoMigrate(&models.User{})
	user.CreatedAt=time.Now();
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": result.Error})
		return
	}

	c.JSON(200, gin.H{"Message": "User Registered","User":user})
}

func Login(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)

	db, err := configs.ConnectDb()
	if err != nil {
		return
	}
	if err := validateEmail(user.UserEmail); !err{
		c.JSON(403, "Invalid email")
		return
	}

	plainText := user.Password

	res := db.Where("user_email= ?", user.UserEmail).First(&user)
	if res.RowsAffected < 1 {
		c.JSON(404, "Email not found")
		return
	}
	isValid := utils.CompareHash(plainText, user.Password)
	if !isValid {
		c.JSON(403, "Invalid password")
		return
	}
	c.JSON(200, "Login Successful")

}
