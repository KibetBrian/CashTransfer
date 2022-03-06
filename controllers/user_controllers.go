package controllers

import (
	"net/http"

	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/models"
	"github.com/KibetBrian/fisa/services"
	"github.com/KibetBrian/fisa/utils"
	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Hello"})
}

//Gets user input, validates and adds to the database
func RegisterUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrResponse(err))
		return
	}
	isValid := utils.CheckValidity(user.Email)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid email"})
		return
	}
	user.Password, _ = utils.HashPassword(user.Password)

	//Check if the user is already registered
	db, err := configs.ConnectDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message":"Error occurred, try again"})
		return
	}
	res := db.Where("email= ?", user.Email).First(&user)
	if res.RowsAffected > 0 {
		c.JSON(http.StatusForbidden, gin.H{"Message":"Email already registered", "Rows Affected": db.RowsAffected})
		return
	}
	registeredUser, err := services.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message":err})
		return
	}
	c.JSON(200, gin.H{"Message": "User Registered", "User": registeredUser})
}

func Login(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)

	db, err := configs.ConnectDb()
	if err != nil {
		return
	}
	if err := utils.CheckValidity(user.Email); !err {
		c.JSON(403, "Invalid email")
		return
	}

	plainText := user.Password
	res := db.Where("user_email= ?", user.Email).First(&user)
	if res.RowsAffected < 1 {
		c.JSON(404, "Email not found")
		return
	}
	isValid := utils.CompareHashAndPassword(user.Password, plainText)
	if !isValid {
		c.JSON(403, "Invalid password")
		return
	}
	c.JSON(200, "Login Successful")
}
